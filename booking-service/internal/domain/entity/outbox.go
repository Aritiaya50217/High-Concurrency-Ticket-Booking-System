package entity

import "time"

type OutboxEvent struct {
	ID          string
	EventType   string
	Payload     string
	Status      string
	CreatedAt   time.Time
	ProcessedAt *time.Time
}
