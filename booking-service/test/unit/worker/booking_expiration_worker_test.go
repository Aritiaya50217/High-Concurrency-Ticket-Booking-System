package worker_test

import (
	"context"
	"testing"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/worker"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/test/unit/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookingExpirationWorker_RunOnce_Success(t *testing.T) {
	bookingRepo := new(mocks.MockBookingRepository)
	outboxRepo := new(mocks.MockOutboxRepository)

	bookingID := uint(1)
	userID := uint(10)
	eventID := uint(100)
	seatID := uint(20)

	booking := &aggregate.Booking{
		ID:      bookingID,
		UserID:  userID,
		EventID: eventID,
		SeatID:  seatID,
		Status:  valueobject.BookingPending,
	}

	bookingRepo.On("FindExpiredBookings", mock.Anything).Return([]*aggregate.Booking{booking}, nil)

	bookingRepo.On("UpdateStatus", mock.Anything, bookingID, "EXPIRED").Return(nil)

	outboxRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	w := worker.NewBookingExpirationWorker(bookingRepo, outboxRepo)

	w.RunOnce(context.Background())

	assert.NotEqual(t, "PENDING", string(booking.Status))

	bookingRepo.AssertCalled(t, "FindExpiredBookings", mock.Anything)

	bookingRepo.AssertCalled(t, "UpdateStatus", mock.Anything, bookingID, "EXPIRED")

	outboxRepo.AssertCalled(t, "Create", mock.Anything, mock.Anything)
}
