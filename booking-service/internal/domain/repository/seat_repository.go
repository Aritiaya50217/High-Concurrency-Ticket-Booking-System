package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
)

type SeatRepository interface {
	FindByID(ctx context.Context, id uint) (*entity.Seat, error)
	LockSeat(ctx context.Context, id uint, version int) error
	MarkBooked(ctx context.Context, id uint) error
}
