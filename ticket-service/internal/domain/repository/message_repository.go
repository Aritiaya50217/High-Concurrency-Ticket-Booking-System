package repository

import "context"

type MessageRepository interface {
	Publish(ctx context.Context, event any) error
}
