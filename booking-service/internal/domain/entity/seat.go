package entity

import "time"

const (
	SeatAvailable string = "AVAILABLE"
	SeatLocked    string = "LOCKED"
	SeatBooked    string = "BOOKED"
)

type Seat struct {
	ID         uint
	EventID    uint
	SeatNumber string
	Status     string
	Version    int // optimistic lock
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
