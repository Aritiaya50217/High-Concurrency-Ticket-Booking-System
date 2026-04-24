package model

import (
	"time"
)

type BookingModel struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"not null;index"`
	EventID   uint   `gorm:"not null;uniqueIndex:idx_event_seat"`
	SeatID    uint   `gorm:"not null;uniqueIndex:idx_event_seat"`
	Status    string `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (BookingModel) TableName() string {
	return "booking"
}
