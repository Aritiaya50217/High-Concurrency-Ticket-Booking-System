package model

import "time"

type SeatModel struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	EventID    uint
	SeatNumber string
	Status     string
	Version    int // optimistic lock
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (SeatModel) TableName() string {
	return "seat"
}
