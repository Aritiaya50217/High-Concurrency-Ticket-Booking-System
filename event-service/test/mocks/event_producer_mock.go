package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockEventProducer struct {
	mock.Mock
}

func (m *MockEventProducer) Publish(ctx context.Context, topic string, message []byte) error {
	args := m.Called(ctx, topic, message)
	return args.Error(0)
}
