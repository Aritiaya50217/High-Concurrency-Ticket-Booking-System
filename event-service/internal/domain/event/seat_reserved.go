package event

import "time"

type SeatReserved struct {
	EventID uint `json:"event_id"`
	SeatID  uint `json:"seat_id"`
	UserID  uint `json:"user_id"`

	OccurredAt time.Time `json:"occurred_at"`
}

func NewSeatReserved(eventID, seatID, userID uint) *SeatReserved {
	return &SeatReserved{EventID: eventID, SeatID: seatID, UserID: userID, OccurredAt: time.Now()}
}
