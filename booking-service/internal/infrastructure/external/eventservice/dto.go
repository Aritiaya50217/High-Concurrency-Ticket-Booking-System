package eventservice

// Anti-Corruption Layer (ACL)
type ReserveSeatRequest struct {
	EventID uint `json:"event_id"`
	SeatID  uint `json:"seat_id"`
}
