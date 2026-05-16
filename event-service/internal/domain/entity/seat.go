package entity

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/valueobject"
)

var (
	ErrSeatNotAvailable = errors.New("seat not available")
)

type Seat struct {
	ID         uint
	EventID    uint
	SeatNumber string
	Status     valueobject.SeatStatus
}

func NewSeat(eventID uint, seatNumber string) *Seat {
	return &Seat{EventID: eventID, SeatNumber: seatNumber}
}

func (s *Seat) Reserve() error {
	if !s.Status.IsAvailable() {
		return ErrSeatNotAvailable
	}

	s.Status = valueobject.SeatReserved
	return nil
}

func (s *Seat) Book() {
	s.Status = valueobject.SeatBooked
}

// ว่าง , cancelled
func (s *Seat) Release() {
	s.Status = valueobject.SeatAvailable
}
