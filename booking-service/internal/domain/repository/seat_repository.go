package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
)

type SeatRepository interface {
	Create(ctx context.Context, seat *entity.Seat) error
	FindByIDForUpdate(ctx context.Context, id uint) (*entity.Seat, error)
	Update(ctx context.Context, seat *entity.Seat) error
}
