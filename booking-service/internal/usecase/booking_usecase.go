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

func (u *BookingUsecase) Update(id uint, status string) (*entity.Booking, error) {
	bookingID, err := u.repo.FindBookingByID(id)
	if err != nil {
		return nil, errors.New("get booking by ID notfound")
	}

	booking := &entity.Booking{
		ID:     bookingID.ID,
		UserID: bookingID.UserID,
		Status: status,
	}

	if err := u.repo.UpdateStatus(booking.ID, booking.Status); err != nil {
		return nil, errors.New("error : update status ")
	}

	return booking, nil
}
