package usecase_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/application/usecase"
// 	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
// 	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"
// 	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
// 	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/test/unit/mocks"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestCreateBookingSuccess(t *testing.T) {
// 	ctx := context.Background()

// 	bookingRepo := new(mocks.MockBookingRepository)
// 	outboxRepo := new(mocks.MockOutboxRepository)

// 	txRepo := mocks.NewTxRepository(bookingRepo, outboxRepo)

// 	bookingRepo.On("WithTransaction", mock.Anything).Return(txRepo, nil)

// 	bookingRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

// 	outboxRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created")

// 	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, booking)

// 	bookingRepo.AssertExpectations(t)
// 	outboxRepo.AssertExpectations(t)
// }

// func TestCreateBookingError(t *testing.T) {
// 	ctx := context.Background()

// 	bookingRepo := new(mocks.MockBookingRepository)
// 	outboxRepo := new(mocks.MockOutboxRepository)

// 	txRepo := mocks.NewTxRepository(bookingRepo, outboxRepo)

// 	bookingRepo.On("WithTransaction", mock.Anything).Return(txRepo, nil)

// 	bookingRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("database error"))

// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "")

// 	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

// 	assert.Error(t, err)
// 	assert.Nil(t, booking)
// }

// func TestCreateBookingOutboxError(t *testing.T) {
// 	ctx := context.Background()

// 	bookingRepo := new(mocks.MockBookingRepository)
// 	outboxRepo := new(mocks.MockOutboxRepository)

// 	txRepo := mocks.NewTxRepository(bookingRepo, outboxRepo)

// 	bookingRepo.On("WithTransaction", mock.Anything).Return(txRepo, nil)

// 	bookingRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

// 	outboxRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("outbox error"))

// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "")

// 	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

// 	assert.Error(t, err)
// 	assert.Nil(t, booking)
// }

// func TestCreateBookingTransactionError(t *testing.T) {
// 	ctx := context.Background()

// 	bookingRepo := new(mocks.MockBookingRepository)

// 	bookingRepo.On("WithTransaction", mock.Anything).Return(nil, errors.New("transaction error"))

// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "")

// 	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

// 	assert.Error(t, err)
// 	assert.Nil(t, booking)
// }

// func TestHandleSeatReservedSuccess(t *testing.T) {
// 	ctx := context.Background()

// 	bookingRepo := new(mocks.MockBookingRepository)

// 	booking := aggregate.NewBooking(1, 100, 10)

// 	bookingRepo.On("FindByEventAndSeat", mock.Anything, uint(100), uint(10)).Return(booking, nil)

// 	bookingRepo.On("UpdateStatus", mock.Anything, booking.ID, string(valueobject.BookingConfirmed)).Return(nil)

// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "")

// 	err := bookingUsecase.HandleSeatReserved(ctx, event.SeatReserved{EventID: 100, SeatID: 10})

// 	assert.NoError(t, err)

// 	bookingRepo.AssertExpectations(t)
// }

// func TestHandleSeatReservedBookingNotFound(t *testing.T) {
// 	ctx := context.Background()

// 	bookingRepo := new(mocks.MockBookingRepository)

// 	bookingRepo.On("FindByEventAndSeat", mock.Anything, uint(100), uint(10)).Return(nil, nil)

// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "")

// 	err := bookingUsecase.HandleSeatReserved(ctx, event.SeatReserved{
// 		EventID: 100,
// 		SeatID:  10,
// 	})

// 	assert.Error(t, err)

// 	assert.Equal(t, "booking not found", err.Error())
// }

// func TestHandleSeatReservedAlreadyConfirmed(t *testing.T) {
// 	ctx := context.Background()

// 	bookingRepo := new(mocks.MockBookingRepository)

// 	booking := aggregate.NewBooking(1, 100, 10)

// 	booking.Confirm()

// 	bookingRepo.On("FindByEventAndSeat", mock.Anything, uint(100), uint(10)).Return(booking, nil)

// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "")

// 	err := bookingUsecase.HandleSeatReserved(ctx, event.SeatReserved{
// 		EventID: 100,
// 		SeatID:  10,
// 	})

// 	assert.NoError(t, err)

// 	bookingRepo.AssertNotCalled(t, "UpdateStatus", mock.Anything, mock.Anything, mock.Anything)
// }

// func TestHandleSeatReservedUpdateStatusError(t *testing.T) {
// 	ctx := context.Background()

// 	bookingRepo := new(mocks.MockBookingRepository)

// 	booking := aggregate.NewBooking(1, 100, 10)

// 	bookingRepo.On("FindByEventAndSeat", mock.Anything, uint(100), uint(10)).Return(booking, nil)

// 	bookingRepo.On("UpdateStatus", mock.Anything, booking.ID, string(valueobject.BookingConfirmed)).Return(errors.New("update error"))

// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "")

// 	err := bookingUsecase.HandleSeatReserved(ctx, event.SeatReserved{
// 		EventID: 100,
// 		SeatID:  10,
// 	})

// 	assert.Error(t, err)

// 	assert.Equal(t, "update error", err.Error())
// }
