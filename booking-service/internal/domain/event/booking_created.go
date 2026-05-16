package event

type BookingCreate struct {
	BookingID uint `json:"booking_id"`
	UserID    uint `json:"user_id"`
	SeatID    uint `json:"seat_id"`
}

func NewBookingCreate(userID, seatID uint) BookingCreate {
	return BookingCreate{
		UserID: userID,
		SeatID: seatID,
	}
}

func (BookingCreate) EventName() string {
	return "BOOKING_CREATED"
}
