package dto

type CreateBookingRequest struct {
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
	SeatID  uint `json:"seat_id"`
}

type UpdateBookingRequest struct {
	Status string `json:"status"`
}

type SearchBookingRequest struct {
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
	Status string `form:"status"`
}
