package repository

import (
	"context"
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
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

func (r *bookingRepository) Create(ctx context.Context, booking *aggregate.Booking) error {
	return r.db.WithContext(ctx).Create(booking).Error
}

func (r *bookingRepository) WithTransaction(ctx context.Context, fn func(repo repository.TxRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := NewTxRepository(
			NewBookingRepository(tx),
			NewOutboxRepository(tx),
		)
		return fn(txRepo)
	})
}

func (r *bookingRepository) FindBySeatID(seatID uint) (*aggregate.Booking, error) {
	var bookingModel model.BookingModel
	err := r.db.Where("seat_id = ?", seatID).First(&bookingModel).Error
	if err != nil {
		return nil, err
	}

	return &aggregate.Booking{}, nil
}

func (r *bookingRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&model.BookingModel{}).Where("id = ?", id).Update("status", status).Error
}

func (r *bookingRepository) FindBookingByID(id uint) (*aggregate.Booking, error) {
	var bookingModel model.BookingModel

	err := r.db.Where("id = ?", id).First(&bookingModel).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("booking not found")
	}
	if err != nil {
		return nil, err
	}

	booking := &aggregate.Booking{
		ID:     bookingModel.ID,
		UserID: bookingModel.UserID,
		SeatID: bookingModel.SeatID,
		// Status:    bookingModel.Status,
		CreatedAt: bookingModel.CreatedAt,
		UpdatedAt: bookingModel.UpdatedAt,
	}

	return booking, nil
}

func (r *bookingRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&model.BookingModel{}).Error
}

func (r *bookingRepository) Search(status string, offset, limit int) ([]aggregate.Booking, int64, error) {
	var bookings []aggregate.Booking
	var models []model.BookingModel
	var total int64

	query := r.db.Model(&model.BookingModel{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	// total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// pagination
	if err := query.Offset(offset).Limit(limit).Find(&models).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range models {
		bookings = append(bookings, aggregate.Booking{
			ID:     m.ID,
			UserID: m.UserID,
			SeatID: m.SeatID,
			// EventID: m.EventID,
			// Status: m.Status,
		})
	}

	return bookings, total, nil

}
