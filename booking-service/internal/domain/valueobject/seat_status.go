package valueobject

type SeatStatus string

const (
	SeatAvailable SeatStatus = "AVAILABLE"
	SeatBooked    SeatStatus = "BOOKED"
)
