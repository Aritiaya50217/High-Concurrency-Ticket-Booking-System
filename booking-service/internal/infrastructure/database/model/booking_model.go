package model

import (
	"time"
)

type BookingModel struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	EventID   uint
	SeatID    uint
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (BookingModel) TableName() string {
	return "booking"
}
