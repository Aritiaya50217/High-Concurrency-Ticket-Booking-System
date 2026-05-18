package event

import "time"

type BookingCreated struct {
	BookingID uint      `json:"booking_id"`
	UserID    uint      `json:"user_id"`
	EventID   uint      `json:"event_id"`
	SeatID    uint      `json:"seat_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewBookingCreated(userID, eventID, seatID uint) BookingCreated {
	return BookingCreated{
		UserID:    userID,
		EventID:   eventID,
		SeatID:    seatID,
		CreatedAt: time.Now(),
	}
}

func (e BookingCreated) EventName() string {
	return "booking.created"
}

func (e BookingCreated) OccurredAt() time.Time {
	return e.CreatedAt
}
