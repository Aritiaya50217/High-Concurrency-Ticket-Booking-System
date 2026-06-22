package repository

import (
	"context"
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	domainRepo "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	valueobject "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
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

func (r *bookingRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&model.BookingModel{}).Where("id = ? AND status = ?", id, "PENDING").Update("status", status).Error
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
		ID:        bookingModel.ID,
		UserID:    bookingModel.UserID,
		SeatID:    bookingModel.SeatID,
		CreatedAt: bookingModel.CreatedAt,
		UpdatedAt: bookingModel.UpdatedAt,
	}

	return booking, nil
}

func (r *bookingRepository) FindByEventAndSeat(ctx context.Context, eventID, seatID uint) (*aggregate.Booking, error) {
	var bookingModel model.BookingModel

	if err := r.db.WithContext(ctx).Where("event_id = ? AND seat_id = ?", eventID, seatID).First(&bookingModel).Error; err != nil {
		return nil, err
	}

	booking := &aggregate.Booking{
		ID:      bookingModel.ID,
		UserID:  bookingModel.UserID,
		EventID: bookingModel.EventID,
		SeatID:  bookingModel.SeatID,
		Status:  valueobject.BookingStatus(bookingModel.Status),
	}
	return booking, nil
}

func (r *bookingRepository) Update(ctx context.Context, booking *aggregate.Booking) error {
	bookingModel := model.BookingModel{
		ID:      booking.ID,
		UserID:  booking.UserID,
		EventID: booking.EventID,
		SeatID:  booking.SeatID,
		Status:  string(booking.Status),
	}

	return r.db.WithContext(ctx).Model(&bookingModel).Where("id = ?", booking.ID).Updates(&bookingModel).Error
}

func (r *bookingRepository) FindExpiredBookings(ctx context.Context) ([]*aggregate.Booking, error) {
	var bookingModels []model.BookingModel
	if err := r.db.WithContext(ctx).Where("status = ?", valueobject.BookingPending).Where("expire_at <= NOW()").Find(&bookingModels).Error; err != nil {
		return nil, err
	}

	bookings := make([]*aggregate.Booking, 0, len(bookingModels))

	for _, m := range bookingModels {
		bookings = append(bookings, &aggregate.Booking{
			ID:       m.ID,
			UserID:   m.UserID,
			EventID:  m.EventID,
			SeatID:   m.SeatID,
			Status:   valueobject.BookingStatus(m.Status),
			ExpireAt: m.ExpireAt,
		})
	}
	return bookings, nil
}
