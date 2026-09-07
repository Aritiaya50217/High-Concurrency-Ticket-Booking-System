package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUser(ctx context.Context, userID uint64) (bool, error) {
	args := m.Called(ctx, userID)

	var exists bool

	if args.Get(0) != nil {
		existsValue := args.Bool(0)
		exists = existsValue
	}

	return exists, args.Error(1)
}
