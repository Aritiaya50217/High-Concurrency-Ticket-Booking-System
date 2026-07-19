package mocks

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockOutboxRepository struct {
	mock.Mock
}

func (m *MockOutboxRepository) Create(ctx context.Context, outbox *entity.OutboxEvent) error {
	args := m.Called(ctx, outbox)

	return args.Error(0)
}

func (m *MockOutboxRepository) FindPending(ctx context.Context) ([]entity.OutboxEvent, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]entity.OutboxEvent), args.Error(1)
}

func (m *MockOutboxRepository) Update(ctx context.Context, outbox *entity.OutboxEvent) error {
	args := m.Called(ctx, outbox)

	return args.Error(0)
}
