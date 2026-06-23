package event

import "time"

type BookingExpired struct {
	BookingID  uint      `json:"booking_id"`
	EventID    uint      `json:"event_id"`
	SeatID     uint      `json:"seat_id"`
	OccurredAt time.Time `json:"occurred_at"`
}

func NewBookingExpired(bookingID, eventID, seatID uint) BookingExpired {
	return BookingExpired{BookingID: bookingID, EventID: eventID, SeatID: seatID}
}
