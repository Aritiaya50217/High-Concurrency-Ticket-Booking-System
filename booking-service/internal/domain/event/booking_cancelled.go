package event

type BookingCancelled struct {
	BookingID uint `json:"booking_id"`
	UserID    uint `json:"user_id"`
	SeatID    uint `json:"seat_id"`
}

func NewBookingCancelled(bookingID, userID, seatID uint) BookingCancelled {
	return BookingCancelled{BookingID: bookingID, UserID: userID, SeatID: seatID}
}

func (BookingCancelled) EventName() string {
	return "BOOKING_CANCELLED"
}
