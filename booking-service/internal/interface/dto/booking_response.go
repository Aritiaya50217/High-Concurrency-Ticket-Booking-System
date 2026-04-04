package dto

type CreateBookingResponse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

type UpdateBookingResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Status string `json:"status"`
}
