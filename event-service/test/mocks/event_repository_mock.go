package mocks

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
	"github.com/stretchr/testify/mock"
)

type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) Create(ctx context.Context, event *aggregate.Event) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *MockEventRepository) Update(ctx context.Context, event *aggregate.Event) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *MockEventRepository) FindByIDForUpdate(ctx context.Context, eventID uint) (*aggregate.Event, error) {
	args := m.Called(ctx, eventID)

	var event *aggregate.Event

	if args.Get(0) != nil {
		event = args.Get(0).(*aggregate.Event)
	}

	return event, args.Error(1)
}

func (m *MockEventRepository) FindByID(ctx context.Context, id uint) (*aggregate.Event, error) {
	args := m.Called(ctx, id)

	var event *aggregate.Event

	if args.Get(0) != nil {
		event = args.Get(0).(*aggregate.Event)
	}

	return event, args.Error(1)
}

func (m *MockEventRepository) Transaction(ctx context.Context, fn func(repository.EventRepository) error) error {
	args := m.Called(ctx, mock.Anything)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return fn(m)
}
