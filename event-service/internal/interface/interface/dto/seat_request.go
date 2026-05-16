package dto

import "time"

type SeatRequest struct {
	// EventID    uint      `json:"event_id"`
	SeatNumber string    `json:"seat_number"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
