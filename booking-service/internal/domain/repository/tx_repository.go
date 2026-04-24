package repository

type TxRepository interface {
	Seat() SeatRepository
	Booking() BookingRepository
	Outbox() OutboxRepository
}
