package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) ReserveSeat(ctx context.Context, eventID, seatID, userID uint) error {
	args := m.Called(ctx, eventID, seatID, userID)
	return args.Error(0)
}
