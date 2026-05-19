package usecase

import (
	"context"
	"log"
	"strconv"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/aggregate"
	domainEvent "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
)

type EventUsecase struct {
	repo          repository.EventRepository
	inboxRepo     repository.InboxRepository
	eventProducer repository.EventProducerRepository
}

func NewEventUsecase(repo repository.EventRepository, inboxRepo repository.InboxRepository, eventProducer repository.EventProducerRepository) *EventUsecase {
	return &EventUsecase{repo: repo, inboxRepo: inboxRepo, eventProducer: eventProducer}
}

func (u *EventUsecase) Create(ctx context.Context, name string) (*aggregate.Event, error) {
	event := aggregate.NewEvent(name)

	if err := u.repo.Create(ctx, event); err != nil {
		return nil, err
	}
	return event, nil

}

func (u *EventUsecase) ReserveSeat(ctx context.Context, eventID, seatID, userID uint) error {
	event, err := u.repo.FindByIDForUpdate(ctx, eventID)
	if err != nil {
		return err
	}

	err = event.ReserveSeat(seatID)
	if err != nil {
		return err
	}

	err = u.repo.Update(ctx, event)
	if err != nil {
		return err
	}

	// publish event (OUTBOX SIDE)
	reservedEvent := domainEvent.NewSeatReserved(eventID, seatID, userID)

	return u.eventProducer.Publish(ctx, reservedEvent)
}

func (u *EventUsecase) HandleBookingCreated(ctx context.Context, event domainEvent.BookingCreated) error {
	// idempotency check
	bookingID := strconv.FormatUint(uint64(event.BookingID), 10)
	processed, err := u.inboxRepo.IsProcessed(ctx, bookingID)
	if err != nil {
		log.Println("IsProcessed error : ", err)
		return err
	}

	if processed {
		return nil
	}

	// business logic
	err = u.repo.Transaction(ctx, func(repo repository.EventRepository) error {
		agg, err := repo.FindByIDForUpdate(ctx, event.EventID)
		if err != nil {
			log.Println("FindByIDForUpdate error : ", err)
			return err
		}

		if err := agg.ReserveSeat(event.SeatID); err != nil {
			log.Println("ReserveSeat error : ", err)
			return err
		}

		// mark processed
		return u.inboxRepo.MarkProcessed(ctx, bookingID, "booking.created")
	})

	return err

}

func (u *EventUsecase) CreateSeats(ctx context.Context, eventID uint, seats []string) error {
	event, err := u.repo.FindByID(ctx, eventID)
	if err != nil {
		return err
	}

	if err = event.CreateSeats(seats); err != nil {
		return err
	}

	return u.repo.Update(ctx, event)

}
