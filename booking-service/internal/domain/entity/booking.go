package entity

import "time"

const (
	StatusPending   string = "PENDING"
	StatusConfirmed string = "CONFIRMED"
	StatusCanceled  string = "CANCELED"
)

type Booking struct {
	ID     uint
	UserID uint
	// EventID   uint
	SeatID    uint
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
