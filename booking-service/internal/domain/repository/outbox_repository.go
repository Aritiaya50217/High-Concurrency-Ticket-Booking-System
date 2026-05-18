package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
)

type OutboxRepository interface {
	Create(ctx context.Context, outbox *entity.OutboxEvent) error
	FindPending(ctx context.Context, limit int) ([]entity.OutboxEvent, error)
	MarkProcessed(ctx context.Context, id uint) error
}
