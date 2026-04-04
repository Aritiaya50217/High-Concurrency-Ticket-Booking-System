package repository

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	domainRepo "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) domainRepo.BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(booking *entity.Booking) error {
	bookingModel := model.BookingModel{
		UserID:  booking.UserID,
		EventID: booking.EventID,
		SeatID:  booking.SeatID,
		Status:  booking.Status,
	}
	err := r.db.Create(&bookingModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *bookingRepository) FindBySeatID(seatID uint) (*entity.Booking, error) {
	var bookingModel model.BookingModel
	err := r.db.Where("seat_id = ?", seatID).First(&bookingModel).Error
	if err != nil {
		return nil, err
	}

	return &entity.Booking{ID: bookingModel.ID, EventID: bookingModel.EventID, SeatID: bookingModel.SeatID, Status: bookingModel.Status}, nil
}

func (r *bookingRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&model.BookingModel{}).Where("id = ?", id).Update("status", status).Error
}

func (r *bookingRepository) FindBookingByID(id uint) (*entity.Booking, error) {
	var bookingModel model.BookingModel

	err := r.db.Where("id = ?", id).First(&bookingModel).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("booking not found")
	}
	if err != nil {
		return nil, err
	}

	booking := &entity.Booking{
		ID:     bookingModel.ID,
		UserID: bookingModel.UserID,
		Status: bookingModel.Status,
	}

	return booking, nil
}
