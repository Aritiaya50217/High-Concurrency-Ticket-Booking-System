package usecase

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
)

type BookingUsecase struct {
	repo repository.BookingRepository
}

func NewBookingUsecase(repo repository.BookingRepository) *BookingUsecase {
	return &BookingUsecase{repo: repo}
}

func (u *BookingUsecase) Create(userID, eventID, seatID uint) (*entity.Booking, error) {
	existing, _ := u.repo.FindBySeatID(seatID)
	if existing != nil && existing.ID != 0 {
		return nil, errors.New("seat already booked")
	}

	booking := &entity.Booking{
		UserID:  userID,
		SeatID:  seatID,
		EventID: eventID,
		Status:  entity.StatusConfirmed,
	}
	if err := u.repo.Create(booking); err != nil {
		return nil, err
	}

	return booking, nil

}
