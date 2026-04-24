package model

import "time"

type OutboxEventModel struct {
	ID          string `gorm:"primaryKey"`
	EventType   string `gorm:"not null"`
	Payload     string `gorm:"type:jsonb"`
	Status      string `gorm:"default:PENDING"`
	CreatedAt   time.Time
	ProcessedAt *time.Time
}

func (OutboxEventModel) TableName() string {
	return "outbox_event"
}
