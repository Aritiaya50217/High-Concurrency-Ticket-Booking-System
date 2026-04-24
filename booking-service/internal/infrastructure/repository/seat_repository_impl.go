package repository

import (
	"context"
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

type seatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) repository.SeatRepository {
	return &seatRepository{db: db}
}

func (r *seatRepository) FindByID(ctx context.Context, id uint) (*entity.Seat, error) {
	var m model.SeatModel
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}

	return &entity.Seat{
		ID:         m.ID,
		EventID:    m.EventID,
		SeatNumber: m.SeatNumber,
		Status:     m.Status,
		Version:    m.Version,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}, nil
}

func (r *seatRepository) LockSeat(ctx context.Context, id uint, version int) error {
	res := r.db.WithContext(ctx).Model(&model.SeatModel{}).Where("id = ? AND version = ? AND status = ?", id, version, entity.SeatAvailable).
		Updates(map[string]interface{}{
			"status":  entity.SeatLocked,
			"version": gorm.Expr("version + 1"),
		})

	if res.RowsAffected == 0 {
		return errors.New("seat already taken")
	}
	return nil
}

func (r *seatRepository) MarkBooked(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.SeatModel{}).Where("id = ?", id).Update("status", entity.SeatBooked).Error
}
