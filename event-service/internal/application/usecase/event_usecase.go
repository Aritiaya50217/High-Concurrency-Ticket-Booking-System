package usecase

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/aggregate"
	domainEvent "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
)

type EventUsecase struct {
	repo        repository.EventRepository
	messageRepo repository.MessageRepository
}

func NewEventUsecase(repo repository.EventRepository, messageRepo repository.MessageRepository) *EventUsecase {
	return &EventUsecase{repo: repo, messageRepo: messageRepo}
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

	reservedEvent := domainEvent.NewSeatReserved(eventID, seatID, userID)

	return u.messageRepo.Publish(ctx, reservedEvent)
}

func (u *EventUsecase) HandleBookingCreated(ctx context.Context, event domainEvent.BookingCreated) error {
	return u.ReserveSeat(ctx, event.EventID, event.SeatID, event.UserID)
}
