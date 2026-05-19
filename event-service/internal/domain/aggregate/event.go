package aggregate

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/entity"
)

var (
	ErrSeatNotFound          = errors.New("seat not found")
	ErrEventCancelled        = errors.New("event cancelled")
	ErrSeatsAlreadyExist     = errors.New("seats already created")
	ErrEventAlreadyCancelled = errors.New("event already cancelled")
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

func (e *Event) CreateSeats(seats []string) error {
	if err := e.isActive(); err != nil {
		return err
	}

	if len(e.Seats) > 0 {
		return ErrSeatsAlreadyExist
	}

	for _, seat := range seats {
		e.Seats = append(e.Seats, entity.NewSeat(e.ID, seat))
	}
	return nil
}

func (e *Event) GetSeat(seatID uint) (*entity.Seat, error) {
	for _, seat := range e.Seats {
		if seat.ID == seatID {
			if seat.Status.IsDeleted() { // seat ที่ถูกลบแล้ว
				return nil, ErrSeatNotFound
			}

			return seat, nil
		}
	}
	return nil, ErrSeatNotFound
}

func (e *Event) isActive() error {
	if e.IsCancelled {
		return ErrEventCancelled
	}
	return nil
}

func (e *Event) ReserveSeat(seatID uint) error {
	if err := e.isActive(); err != nil {
		return err
	}

	seat, err := e.GetSeat(seatID)
	if err != nil {
		return err
	}

	return seat.Reserve()
}

func (e *Event) ReleaseSeat(seatID uint) error {
	if err := e.isActive(); err != nil {
		return err
	}

	seat, err := e.GetSeat(seatID)
	if err != nil {
		return err
	}
	return seat.Release()
}