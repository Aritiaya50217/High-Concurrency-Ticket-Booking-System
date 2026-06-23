package entity

import "time"

const (
	OutboxPending = "PENDING"
	OutboxSent    = "SENT"
)

type OutboxEvent struct {
	ID          uint
	EventType   string
	Payload     string
	Status      string
	CreatedAt   time.Time
	ProcessedAt time.Time
	MessageID   string
}
