package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/external/eventservice"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/test/unit/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBookingSuccess(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	outboxRepo := new(mocks.MockOutboxRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	txRepo := mocks.NewTxRepository(bookingRepo, outboxRepo)

	exists := true

	userService.On("GetUser", mock.Anything, mock.AnythingOfType("uint64")).Return(exists, nil)

	bookingRepo.On("WithTransaction", mock.Anything).Return(txRepo, nil)

	bookingRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	outboxRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

	assert.NoError(t, err)
	assert.NotNil(t, booking)

	userService.AssertExpectations(t)
	bookingRepo.AssertExpectations(t)
	outboxRepo.AssertExpectations(t)
}

func TestCreateBookingError(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	outboxRepo := new(mocks.MockOutboxRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	txRepo := mocks.NewTxRepository(bookingRepo, outboxRepo)

	exists := true

	userService.On("GetUser", mock.Anything, mock.AnythingOfType("uint64")).Return(exists, nil)

	bookingRepo.On("WithTransaction", mock.Anything).Return(txRepo, nil)

	bookingRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("database error"))

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

	assert.Error(t, err)
	assert.Nil(t, booking)

	userService.AssertExpectations(t)
	bookingRepo.AssertExpectations(t)
	outboxRepo.AssertExpectations(t)
}

func TestCreateBookingOutboxError(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	outboxRepo := new(mocks.MockOutboxRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	txRepo := mocks.NewTxRepository(bookingRepo, outboxRepo)

	exists := true

	userService.On("GetUser", mock.Anything, mock.AnythingOfType("uint64")).Return(exists, nil)

	bookingRepo.On("WithTransaction", mock.Anything).Return(txRepo, nil)

	bookingRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	outboxRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("outbox error"))

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

	assert.Error(t, err)
	assert.Nil(t, booking)

	userService.AssertExpectations(t)
	bookingRepo.AssertExpectations(t)
	outboxRepo.AssertExpectations(t)
}

func TestCreateBookingTransactionError(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	exists := true

	userService.On("GetUser", mock.Anything, mock.AnythingOfType("uint64")).Return(exists, nil)

	bookingRepo.On("WithTransaction", mock.Anything).Return(nil, errors.New("transaction error"))

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

	assert.Error(t, err)
	assert.Nil(t, booking)

	userService.AssertExpectations(t)
	bookingRepo.AssertExpectations(t)
}

func TestHandleSeatReservedSuccess(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	exists := true

	booking := aggregate.NewBooking(1, 100, 10)

	userService.On("GetUser", mock.Anything, mock.AnythingOfType("uint64")).Return(exists, nil)

	bookingRepo.On("FindByEventAndSeat", mock.Anything, uint(100), uint(10)).Return(booking, nil)

	bookingRepo.On("UpdateStatus", mock.Anything, booking.ID, string(valueobject.BookingConfirmed)).Return(nil)

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	err := bookingUsecase.HandleSeatReserved(ctx, event.SeatReserved{EventID: 100, SeatID: 10})

	assert.NoError(t, err)

	bookingRepo.AssertExpectations(t)
}

func TestHandleSeatReservedBookingNotFound(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	bookingRepo.On("FindByEventAndSeat", mock.Anything, uint(100), uint(10)).Return(nil, nil)

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	err := bookingUsecase.HandleSeatReserved(ctx, event.SeatReserved{
		EventID: 100,
		SeatID:  10,
	})

	assert.Error(t, err)

	assert.Equal(t, "booking not found", err.Error())
}

func TestHandleSeatReservedAlreadyConfirmed(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	booking := aggregate.NewBooking(1, 100, 10)

	booking.Confirm()

	bookingRepo.On("FindByEventAndSeat", mock.Anything, uint(100), uint(10)).Return(booking, nil)

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	err := bookingUsecase.HandleSeatReserved(ctx, event.SeatReserved{
		EventID: 100,
		SeatID:  10,
	})

	assert.NoError(t, err)

	bookingRepo.AssertNotCalled(t, "UpdateStatus", mock.Anything, mock.Anything, mock.Anything)
}

func TestHandleSeatReservedUpdateStatusError(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	booking := aggregate.NewBooking(1, 100, 10)

	bookingRepo.On("FindByEventAndSeat", mock.Anything, uint(100), uint(10)).Return(booking, nil)

	bookingRepo.On("UpdateStatus", mock.Anything, booking.ID, string(valueobject.BookingConfirmed)).Return(errors.New("update error"))

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	err := bookingUsecase.HandleSeatReserved(ctx, event.SeatReserved{
		EventID: 100,
		SeatID:  10,
	})

	assert.Error(t, err)

	assert.Equal(t, "update error", err.Error())
}

func TestCreateBookingUserExists(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	outboxRepo := new(mocks.MockOutboxRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	txRepo := mocks.NewTxRepository(bookingRepo, outboxRepo)

	exists := true

	userService.On("GetUser", mock.Anything, mock.AnythingOfType("uint64")).Return(exists, nil)

	bookingRepo.On("WithTransaction", mock.Anything).Return(txRepo, nil)

	bookingRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	outboxRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

	assert.NoError(t, err)
	assert.NotNil(t, booking)

	userService.AssertExpectations(t)
	bookingRepo.AssertExpectations(t)
	outboxRepo.AssertExpectations(t)
}

func TestCreateBookingUserNotFound(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	exists := false

	userService.On("GetUser", mock.Anything, mock.AnythingOfType("uint64")).Return(exists, nil)

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

	assert.Error(t, err)
	assert.Nil(t, booking)

	userService.AssertExpectations(t)

	// user ไม่พบ -> ไม่ควรแตะ DB transaction
	bookingRepo.AssertNotCalled(t, "WithTransaction", mock.Anything)
}

func TestCreateBookingUserServiceError(t *testing.T) {
	ctx := context.Background()

	bookingRepo := new(mocks.MockBookingRepository)
	userService := new(mocks.MockUserService)
	eventClient := eventservice.SeatServiceClient{}

	userService.On("GetUser", mock.Anything, mock.AnythingOfType("uint64")).Return(false, errors.New("user service unavailable"))

	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, nil, "booking.created", eventClient, userService)

	booking, err := bookingUsecase.Create(ctx, 1, 100, 10)

	assert.Error(t, err)
	assert.Nil(t, booking)

	bookingRepo.AssertNotCalled(t, "WithTransaction", mock.Anything)
}
