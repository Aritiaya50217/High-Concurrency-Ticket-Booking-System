package dto

type CreateBookingRequest struct {
	EventID uint `json:"event_id"`
	SeatID  uint `json:"seat_id"`
}

type UpdateBookingRequest struct {
	Status string `json:"status"`
}