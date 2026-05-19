package dto

type SeatReservedEvent struct {
	BookingID uint `json:"booking_id"`
	SeatID    uint `json:"seat_id"`
	EventID   uint `json:"event_id"`
}
