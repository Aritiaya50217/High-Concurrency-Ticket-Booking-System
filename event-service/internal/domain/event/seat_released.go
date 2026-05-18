package event

import "time"

type SeatReleased struct {
	EventID uint `json:"event_id"`
	SeatID  uint `json:"seat_id"`

	OccurredAt time.Time `json:"occurred_at"`
}

func NewSeatReleased(eventID, seatID uint) *SeatReleased {

	return &SeatReleased{EventID: eventID, SeatID: seatID, OccurredAt: time.Now()}
}
