package aggregate

import (
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
		Status:  valueobject.BookingConfirmed,
		events:  []event.DomainEvent{},
	}

	booking.addEvent(event.NewBookingCreated(userID, eventID, seatID))

	return booking

}

// func (b *Booking) Cancel() {
// 	b.Status = valueobject.BookingCancelled

// 	b.addEvent(event.NewBookingCancelled(b.ID, b.UserID, b.SeatID))
// }

func (b *Booking) Events() []event.DomainEvent {
	return b.events
}

func (b *Booking) ClearEvents() {
	b.events = nil
}

func (b *Booking) addEvent(e event.DomainEvent) {
	b.events = append(b.events, e)
}
