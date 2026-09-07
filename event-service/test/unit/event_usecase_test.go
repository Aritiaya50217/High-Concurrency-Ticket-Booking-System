package unit

import (
	"context"
	"errors"
	"testing"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/valueobject"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateEventSuccess(t *testing.T) {
	repo := new(mocks.MockEventRepository)
	producer := new(mocks.MockEventProducer)

	name := "Concert"

	repo.On("Create", mock.Anything, mock.AnythingOfType("*aggregate.Event")).Return(nil)

	u := usecase.NewEventUsecase(repo, nil, producer)

	event, err := u.Create(context.Background(), name)

	require.NoError(t, err)
	require.NotNil(t, event)
	require.Equal(t, name, event.Name)

	repo.AssertExpectations(t)
}

func TestCreateEventRepositoryError(t *testing.T) {
	repo := new(mocks.MockEventRepository)
	producer := new(mocks.MockEventProducer)

	expectedErr := errors.New("database error")
	name := "Concert"

	repo.On("Create", mock.Anything, mock.AnythingOfType("*aggregate.Event")).Return(expectedErr)

	u := usecase.NewEventUsecase(repo, nil, producer)

	event, err := u.Create(context.Background(), name)

	require.ErrorIs(t, err, expectedErr)
	require.Nil(t, event)

	repo.AssertExpectations(t)
}

func TestReserveSeatSuccess(t *testing.T) {
	repo := new(mocks.MockEventRepository)
	producer := new(mocks.MockEventProducer)

	eventID := uint(1)
	seatID := uint(10)
	userID := uint(100)

	seatNumber := "A1"

	event := &aggregate.Event{
		ID:          eventID,
		Name:        "Concert",
		IsCancelled: false,
		Seats: []*entity.Seat{
			{
				ID:         seatID,
				EventID:    eventID,
				SeatNumber: seatNumber,
				Status:     valueobject.SeatAvailable,
				Version:    0,
			},
		},
	}

	repo.On("FindByIDForUpdate", mock.Anything, eventID).Return(event, nil)

	repo.On("Update", mock.Anything, event).Return(nil)

	producer.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	u := usecase.NewEventUsecase(repo, nil, producer)

	err := u.ReserveSeat(context.Background(), eventID, seatID, userID)

	require.NoError(t, err)

	require.Equal(t, valueobject.SeatReserved, event.Seats[0].Status)

	repo.AssertExpectations(t)
	producer.AssertExpectations(t)
}

func TestReserveSeatEventNotFound(t *testing.T) {
	repo := new(mocks.MockEventRepository)
	producer := new(mocks.MockEventProducer)

	eventID := uint(1)
	seatID := uint(10)
	userID := uint(100)

	expectedErr := errors.New("event not found")

	repo.On("FindByIDForUpdate", mock.Anything, eventID).Return(nil, expectedErr)

	u := usecase.NewEventUsecase(repo, nil, producer)

	err := u.ReserveSeat(context.Background(), eventID, seatID, userID)

	require.ErrorIs(t, err, expectedErr)

	repo.AssertExpectations(t)

	producer.AssertNotCalled(t, "Publish", mock.Anything, mock.Anything, mock.Anything)
}

func TestReserveSeatSeatNotFound(t *testing.T) {
	repo := new(mocks.MockEventRepository)
	producer := new(mocks.MockEventProducer)

	eventID := uint(1)
	seatID := uint(999)
	userID := uint(100)

	seatNumber := "A1"

	event := &aggregate.Event{
		ID:          eventID,
		Name:        "Concert",
		IsCancelled: false,
		Seats: []*entity.Seat{
			{
				ID:         10,
				EventID:    eventID,
				SeatNumber: seatNumber,
				Status:     valueobject.SeatAvailable,
				Version:    0,
			},
		},
	}

	repo.On("FindByIDForUpdate", mock.Anything, eventID).Return(event, nil)

	u := usecase.NewEventUsecase(repo, nil, producer)

	err := u.ReserveSeat(context.Background(), eventID, seatID, userID)

	require.ErrorIs(t, err, aggregate.ErrSeatNotFound)

	repo.AssertExpectations(t)
	repo.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
	producer.AssertNotCalled(t, "Publish", mock.Anything, mock.Anything, mock.Anything)
}

func TestReserveSeatAlreadyReserved(t *testing.T) {
	repo := new(mocks.MockEventRepository)
	producer := new(mocks.MockEventProducer)

	eventID := uint(1)
	seatID := uint(10)
	userID := uint(100)
	seatNumber := "A1"

	event := &aggregate.Event{
		ID:          eventID,
		Name:        "Concert",
		IsCancelled: false,
		Seats: []*entity.Seat{
			{
				ID:         seatID,
				EventID:    eventID,
				SeatNumber: seatNumber,
				Status:     valueobject.SeatReserved,
				Version:    0,
			},
		},
	}

	repo.On("FindByIDForUpdate", mock.Anything, eventID).Return(event, nil)

	u := usecase.NewEventUsecase(repo, nil, producer)

	err := u.ReserveSeat(context.Background(), eventID, seatID, userID)

	require.ErrorIs(t, err, entity.ErrSeatNotAvailable)

	repo.AssertExpectations(t)
	repo.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
	producer.AssertNotCalled(t, "Publish", mock.Anything, mock.Anything, mock.Anything)
}

func TestReserveSeatUpdateError(t *testing.T) {
	repo := new(mocks.MockEventRepository)
	producer := new(mocks.MockEventProducer)

	eventID := uint(1)
	seatID := uint(10)
	userID := uint(100)
	seatNumber := "A1"

	expectedErr := errors.New("update failed")

	event := &aggregate.Event{
		ID:          eventID,
		Name:        "Concert",
		IsCancelled: false,
		Seats: []*entity.Seat{
			{
				ID:         seatID,
				EventID:    eventID,
				SeatNumber: seatNumber,
				Status:     valueobject.SeatAvailable,
				Version:    0,
			},
		},
	}

	repo.On("FindByIDForUpdate", mock.Anything, eventID).Return(event, nil)

	repo.On("Update", mock.Anything, event).Return(expectedErr)

	u := usecase.NewEventUsecase(repo, nil, producer)

	err := u.ReserveSeat(context.Background(), eventID, seatID, userID)

	require.ErrorIs(t, err, expectedErr)

	repo.AssertExpectations(t)

	producer.AssertNotCalled(t, "Publish", mock.Anything, mock.Anything, mock.Anything)
}
