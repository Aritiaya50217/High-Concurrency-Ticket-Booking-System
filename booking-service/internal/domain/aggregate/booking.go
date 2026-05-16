package aggregate

import (
	"errors"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
)

type Booking struct {
	ID     uint
	UserID uint
	SeatID uint
	Status valueobject.BookingStatus

	// Domain Events
	events    []event.DomainEvent
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBooking(userID uint, seat *entity.Seat) (*Booking, error) {
	// business rule
	if seat == nil {
		return nil, errors.New("seat not found")
	}

	if err := seat.Book(); err != nil {
		return nil, err
	}

	booking := &Booking{
		UserID: userID,
		SeatID: seat.ID,
		Status: valueobject.BookingConfirmed,
	}

	booking.addEvent(event.NewBookingCreate(userID, seat.ID))

	return booking, nil

}

func (b *Booking) Cancel() {
	b.Status = valueobject.BookingCancelled

	b.addEvent(event.NewBookingCancelled(b.ID, b.UserID, b.SeatID))
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
