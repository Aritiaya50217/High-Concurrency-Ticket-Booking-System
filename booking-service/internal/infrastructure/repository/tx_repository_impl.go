package repository

import "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"

type txRepository struct {
	// seatRepo    repository.SeatRepository
	bookingRepo repository.BookingRepository
	outboxRepo  repository.OutboxRepository
}

// func (t *txRepository) Seat() repository.SeatRepository {
// 	return t.seatRepo
// }

func (t *txRepository) Booking() repository.BookingRepository {
	return t.bookingRepo
}

func (t *txRepository) Outbox() repository.OutboxRepository {
	return t.outboxRepo
}
