package dto

type CreateBookingRequest struct {
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
	SeatID  uint `json:"seat_id"`
}
