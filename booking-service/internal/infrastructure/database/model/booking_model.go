package model

import (
	"time"
)

type BookingModel struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"not null;index"`
	EventID   uint   `gorm:"not null;index"`
	SeatID    uint   `gorm:"not null;index"`
	Status    string `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpireAt  time.Time
}

func (BookingModel) TableName() string {
	return "bookings"
}
