package dto

type BookingCreatedEvent struct {
	Event     string `json:"event"`
	BookingID uint   `json:"booking_id"`
	UserID    uint   `json:"user_id"`
	// EventID   uint   `json:"event_id"`
	SeatID uint `json:"seat_id"`
}
