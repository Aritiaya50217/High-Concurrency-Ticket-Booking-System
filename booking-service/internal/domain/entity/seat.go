package entity

import (
	"errors"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
)

type Seat struct {
	ID         uint
	SeatNumber string
	Status     string
	Version    int // optimistic lock
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (s *Seat) Book() error {
	if s.Status != string(valueobject.SeatAvailable) {
		return errors.New("seat already booked")
	}

	s.Status = string(valueobject.SeatBooked)

	return nil
}
