package event

import "time"

type BookingCancelled struct {
	BookingID uint `json:"booking_id"`
	EventID   uint `json:"event_id"`
	SeatID    uint `json:"seat_id"`

	OccurredAt time.Time `json:"occurred_at"`
}

func NewBookingCancelled(bookingID uint, eventID uint, seatID uint) *BookingCancelled {
	return &BookingCancelled{
		BookingID:  bookingID,
		EventID:    eventID,
		SeatID:     seatID,
		OccurredAt: time.Now(),
	}
}
