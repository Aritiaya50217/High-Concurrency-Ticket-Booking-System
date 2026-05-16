package valueobject

type SeatStatus string

const (
	SeatAvailable SeatStatus = "AVAILABLE"
	SeatReserved  SeatStatus = "RESERVED"
	SeatBooked    SeatStatus = "BOOKED"
)

func (s SeatStatus) IsAvailable() bool {
	return s == SeatAvailable
}

func (s SeatStatus) IsReserved() bool {
	return s == SeatReserved
}

func (s SeatStatus) IsBooked() bool {
	return s == SeatBooked
}
