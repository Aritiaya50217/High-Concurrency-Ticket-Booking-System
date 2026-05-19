package event

type SeatReserved struct {
	BookingID uint
	EventID   uint
	SeatID    uint
	UserID    uint
}
