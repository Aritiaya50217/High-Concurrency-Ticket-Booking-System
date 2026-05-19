package event

import "time"

type BookingConfirmed struct {
	BookingID uint      `json:"booking_id"`
	UserID    uint      `json:"user_id"`
	EventID   uint      `json:"event_id"`
	SeatID    uint      `json:"seat_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewBookingConfirmed(bookingID, userID, eventID, seatID uint) BookingConfirmed {
	return BookingConfirmed{
		BookingID: bookingID,
		UserID:    userID,
		EventID:   eventID,
		SeatID:    seatID,
		CreatedAt: time.Now(),
	}
}

func (e BookingConfirmed) EventName() string {
	return "booking.confirmed"
}

func (e BookingConfirmed) OccurredAt() time.Time {
	return e.CreatedAt
}
