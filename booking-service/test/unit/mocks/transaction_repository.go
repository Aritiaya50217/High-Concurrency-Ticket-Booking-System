package mocks

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
)

type MockTxRepository struct {
	BookingRepo repository.BookingRepository
	OutboxRepo  repository.OutboxRepository
}

func NewTxRepository(bookingRepo repository.BookingRepository,outboxRepo repository.OutboxRepository) *MockTxRepository {
	return &MockTxRepository{
		BookingRepo: bookingRepo,
		OutboxRepo:  outboxRepo,
	}
}

func (m *MockTxRepository) Booking() repository.BookingRepository {
	return m.BookingRepo
}

func (m *MockTxRepository) Outbox() repository.OutboxRepository {
	return m.OutboxRepo
}
