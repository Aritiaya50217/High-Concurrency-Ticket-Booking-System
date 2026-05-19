package model

import "time"

type ProcessedEventModel struct {
	ID        uint `gorm:"primaryKey"`
	EventID   string
	EventType string
	CreatedAt time.Time
}

func (ProcessedEventModel) TableName() string {
	return "processed_events"
}
