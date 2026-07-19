package mocks

import (
	"context"
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/stretchr/testify/mock"
)

type MockBookingRepository struct {
	mock.Mock
}

func (m *MockBookingRepository) Create(ctx context.Context, booking *aggregate.Booking) error {
	args := m.Called(ctx, booking)
	return args.Error(0)
}

func (m *MockBookingRepository) Update(ctx context.Context, booking *aggregate.Booking) error {
	args := m.Called(ctx, booking)
	return args.Error(0)
}

func (m *MockBookingRepository) WithTransaction(ctx context.Context, fn func(repo repository.TxRepository) error) error {
	args := m.Called(ctx)

	if args.Error(1) != nil {
		return args.Error(1)
	}

	txRepo, ok := args.Get(0).(repository.TxRepository)
	if !ok {
		return errors.New("invalid tx repository")
	}

	return fn(txRepo)
}

func (m *MockBookingRepository) FindBySeatID(seatID uint) (*aggregate.Booking, error) {
	args := m.Called(seatID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*aggregate.Booking), args.Error(1)
}

func (m *MockBookingRepository) FindBookingByID(id uint) (*aggregate.Booking, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*aggregate.Booking), args.Error(1)
}

func (m *MockBookingRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

func (m *MockBookingRepository) FindByEventAndSeat(ctx context.Context, eventID, seatID uint) (*aggregate.Booking, error) {
	args := m.Called(ctx, eventID, seatID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*aggregate.Booking), args.Error(1)
}

func (m *MockBookingRepository) FindExpiredBookings(ctx context.Context) ([]*aggregate.Booking, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*aggregate.Booking), args.Error(1)
}
