package event

type BookingCreated struct {
	BookingID uint `json:"booking_id"`
	UserID    uint `json:"user_id"`
	EventID   uint `json:"event_id"`
	SeatID    uint `json:"seat_id"`
}

func NewBookingCreated(userID, eventID, seatID uint) BookingCreated {
	return BookingCreated{
		UserID:  userID,
		EventID: eventID,
		SeatID:  seatID,
	}
}

func (BookingCreated) EventName() string {
	return "BOOKING_CREATED"
}
