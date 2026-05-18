package dto

type CreateSeatRequest struct {
	EventID uint     `json:"event_id" binding:"required"`
	Seats   []string `json:"seats" binding:"required"`
}

type ReserveSeatRequest struct {
	EventID uint `json:"event_id" binding:"required"`
	SeatID  uint `json:"seat_id" binding:"required"`
}

type ReleaseSeatRequest struct {
	EventID uint `json:"event_id" binding:"required"`
	SeatID  uint `json:"seat_id" binding:"required"`
}
