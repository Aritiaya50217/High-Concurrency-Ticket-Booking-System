package usecase

import (
	"context"
	"encoding/json"
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

	if err = event.ReserveSeat(seatID); err != nil {
		return err
	}

	if err := u.repo.Update(ctx, event); err != nil {
		return err
	}

	// publish event
	reservedEvent := domainEvent.NewSeatReserved(eventID, seatID, userID)

	data, err := json.Marshal(reservedEvent)
	if err != nil {
		log.Println("SeatReserved error : ", err)
		return err
	}
	return u.eventProducer.Publish(ctx, "seat.reserved", data)
}

func (u *EventUsecase) HandleBookingCreated(ctx context.Context, event domainEvent.BookingCreated) error {
	// idempotency check
	eventID := strconv.FormatUint(uint64(event.EventID), 10)

	processed, err := u.inboxRepo.IsProcessed(ctx, eventID)
	if err != nil {
		log.Println("IsProcessed error : ", err)
		return err
	}

	if processed {
		return nil
	}

	// booking.created is handled asynchronously.
	// Seat reservation is handled by gRPC ReserveSeat()

	if err := u.inboxRepo.MarkProcessed(
		ctx,
		eventID,
		"booking.created",
	); err != nil {
		log.Println("MarkProcessed error : ", err)
		return err
	}

	return nil

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

func (u *EventUsecase) HandleBookingCancelled(ctx context.Context, event domainEvent.BookingCancelled) error {
	return u.handleBookingCancelled(ctx, event)
}

func (u *EventUsecase) handleBookingCancelled(ctx context.Context, event domainEvent.BookingCancelled) error {
	log.Println("releasing seat:", event.SeatID)
	return u.repo.Transaction(ctx, func(repo repository.EventRepository) error {
		eventAgg, err := repo.FindByIDForUpdate(ctx, event.EventID)
		if err != nil {
			return err
		}

		if err := eventAgg.ReleaseSeat(event.SeatID); err != nil {
			return err
		}
		return repo.Update(ctx, eventAgg)
	})
}
