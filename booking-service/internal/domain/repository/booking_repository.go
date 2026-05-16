package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *aggregate.Booking) error
	WithTransaction(ctx context.Context, fn func(repo TxRepository) error) error
	FindBySeatID(seatID uint) (*aggregate.Booking, error)
	FindBookingByID(id uint) (*aggregate.Booking, error)
	UpdateStatus(id uint, status string) error
	Delete(id uint) error
	Search(status string, offset, limit int) ([]aggregate.Booking, int64, error)
}
