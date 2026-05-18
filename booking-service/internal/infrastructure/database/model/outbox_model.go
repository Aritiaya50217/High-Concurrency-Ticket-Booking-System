package model

import (
	"time"

	"gorm.io/datatypes"
)

type OutboxEventModel struct {
	ID          uint           `gorm:"primaryKey"`
	EventType   string         `gorm:"not null"`
	Payload     datatypes.JSON `gorm:"type:jsonb"`
	Status      string         `gorm:"default:PENDING"`
	CreatedAt   time.Time
	ProcessedAt time.Time
}

func (OutboxEventModel) TableName() string {
	return "outbox_events"
}
