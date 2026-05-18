package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
)

type OutboxRepository interface {
	Create(ctx context.Context, outbox *entity.OutboxEvent) error
	FindPending(ctx context.Context) ([]entity.OutboxEvent, error)
	Update(ctx context.Context, outbox *entity.OutboxEvent) error
}
