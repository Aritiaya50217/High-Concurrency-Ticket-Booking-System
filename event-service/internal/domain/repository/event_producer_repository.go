package repository

import (
	"context"
)

type EventProducerRepository interface {
	Publish(ctx context.Context, event any) error
}