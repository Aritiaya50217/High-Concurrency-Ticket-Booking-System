package repository

import "context"

// consumer
type InboxRepository interface {
	IsProcessed(ctx context.Context, eventID string) (bool, error)
	MarkProcessed(ctx context.Context, eventID string, eventType string) error
}
