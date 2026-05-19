package repository

import (
	"context"
)

type EventProducerRepository interface {
	Publish(ctx context.Context, topic string, message []byte) error
}
