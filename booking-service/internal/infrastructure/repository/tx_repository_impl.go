package repository

import "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"

type txRepository struct {
	bookingRepo repository.BookingRepository
	outboxRepo  repository.OutboxRepository
}

func NewTxRepository(bookingRepo repository.BookingRepository, outboxRepo repository.OutboxRepository) repository.TxRepository {

	return &txRepository{
		bookingRepo: bookingRepo,
		outboxRepo:  outboxRepo,
	}
}

func (t *txRepository) Booking() repository.BookingRepository {
	return t.bookingRepo
}

func (t *txRepository) Outbox() repository.OutboxRepository {
	return t.outboxRepo
}
