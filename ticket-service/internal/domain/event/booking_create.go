package event

import "time"

type BookingCreated struct {
	BookingID uint      `json:"booking_id"`
	EventID   uint      `json:"event_id"`
	SeatID    uint      `json:"seat_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewBookingCreated(bookingID, eventID, seatID, userID uint) *BookingCreated {
	return &BookingCreated{BookingID: bookingID, EventID: eventID, SeatID: seatID, UserID: userID, CreatedAt: time.Now()}
}
