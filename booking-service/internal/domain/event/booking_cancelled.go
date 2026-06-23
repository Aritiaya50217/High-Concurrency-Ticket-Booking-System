package event

type BookingCancelled struct {
	BookingID uint `json:"booking_id"`
	UserID    uint `json:"user_id"`
	SeatID    uint `json:"seat_id"`
	EventID   uint `json:"event_id"`
}

func NewBookingCancelled(bookingID, userID, seatID, eventID uint) BookingCancelled {
	return BookingCancelled{BookingID: bookingID, UserID: userID, SeatID: seatID, EventID: eventID}
}

func (BookingCancelled) EventName() string {
	return "booking.cancelled"
}
