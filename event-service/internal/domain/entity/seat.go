package entity

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/valueobject"
)

var (
	ErrSeatNotAvailable = errors.New("seat not available")
	ErrSeatReserved     = errors.New("seat must be reserved before booking")
	ErrSeatNotReserved  = errors.New("seat is not reserved")
	ErrSeatBooked       = errors.New("cannot release booked seat")
	ErrSeatDeleted      = errors.New("cannot delete booked seat")
)

type Seat struct {
	ID         uint
	EventID    uint
	SeatNumber string
	Status     valueobject.SeatStatus
	Version    int
}

func NewSeat(eventID uint, seatNumber string) *Seat {
	return &Seat{EventID: eventID, SeatNumber: seatNumber, Status: valueobject.SeatAvailable, Version: 0}
}

func (s *Seat) Reserve() error {
	if !s.Status.IsAvailable() {
		return ErrSeatNotAvailable
	}

	s.Status = valueobject.SeatReserved
	return nil
}

func (s *Seat) Release() error {
	if !s.Status.IsAvailable() {
		return ErrSeatNotAvailable
	}

	s.Status = valueobject.SeatAvailable
	return nil
}
