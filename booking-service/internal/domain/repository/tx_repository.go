package repository

type TxRepository interface {
	Booking() BookingRepository
	Outbox() OutboxRepository
}
