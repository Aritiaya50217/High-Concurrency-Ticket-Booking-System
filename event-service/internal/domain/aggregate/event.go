package aggregate

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/entity"
)

var (
	ErrSeatNotFound   = errors.New("seat not found")
	ErrEventCancelled = errors.New("event cancelled")
)

type Event struct {
	ID          uint
	Name        string
	IsCancelled bool
	Seats       []*entity.Seat
}

func NewEvent(name string) *Event {
	return &Event{Name: name, IsCancelled: false, Seats: []*entity.Seat{}}
}

func (e *Event) AddSeat(number string) {
	seat := entity.NewSeat(e.ID, number)

	e.Seats = append(e.Seats, seat)

}

func (e *Event) Cancel() {
	e.IsCancelled = true
}

func (e *Event) ReserveSeat(seatID uint) error {
	if e.IsCancelled {
		return ErrEventCancelled
	}

	for _, seat := range e.Seats {
		if seat.ID == seatID {
			return seat.Reserve()
		}
	}
	return ErrSeatNotFound
}

func (e *Event) ReleaseSeat(seatID uint) error {
	for _, seat := range e.Seats {
		if seat.ID != seatID {
			continue
		}
		seat.Release()

		return nil
	}
	return ErrSeatNotFound
}

func (e *Event) BookSeat(seatID uint) error {
	for _, seat := range e.Seats {
		if seat.ID != seatID {
			continue
		}
		seat.Book()
		return nil
	}

	return ErrSeatNotFound
}
