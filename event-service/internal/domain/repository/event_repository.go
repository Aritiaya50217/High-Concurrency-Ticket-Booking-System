package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/aggregate"
)

type EventRepository interface {
	Create(ctx context.Context, event *aggregate.Event) error
	WithTransaction(ctx context.Context, fn func(repo EventRepository) error) error
	Update(ctx context.Context, event *aggregate.Event) error
	FindByIDForUpdate(ctx context.Context, eventID uint) (*aggregate.Event, error)
}
