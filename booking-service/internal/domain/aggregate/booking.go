package aggregate

import (
	"errors"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
)

type Booking struct {
	ID      uint
	UserID  uint
	EventID uint
	SeatID  uint
	Status  valueobject.BookingStatus

	// Domain Events
	events    []event.DomainEvent
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBooking(userID uint, eventID uint, seatID uint) *Booking {
	// business rule
	booking := &Booking{
		UserID:  userID,
		EventID: eventID,
		SeatID:  seatID,
		Status:  valueobject.BookingPending,
		events:  []event.DomainEvent{},
	}

	booking.addEvent(event.NewBookingCreated(userID, eventID, seatID))

	return booking

}

func (b *Booking) Events() []event.DomainEvent {
	return b.events
}

func (b *Booking) ClearEvents() {
	b.events = nil
}

func (b *Booking) addEvent(e event.DomainEvent) {
	b.events = append(b.events, e)
}

func (b *Booking) Confirm() error {
	if b.Status == valueobject.BookingConfirmed {
		return nil
	}

	if b.Status != valueobject.BookingPending {
		return errors.New("booking must be PENDING to confirm")
	}

	b.Status = valueobject.BookingConfirmed

	b.addEvent(event.NewBookingConfirmed(b.ID, b.UserID, b.EventID, b.SeatID))

	return nil
}
