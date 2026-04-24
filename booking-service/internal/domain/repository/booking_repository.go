package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *entity.Booking) error
	WithTransaction(ctx context.Context, fn func(repo TxRepository) error) error
	FindBySeatID(seatID uint) (*entity.Booking, error)
	FindBookingByID(id uint) (*entity.Booking, error)
	UpdateStatus(id uint, status string) error
	Delete(id uint) error
	Search(status string, offset, limit int) ([]entity.Booking, int64, error)
}
