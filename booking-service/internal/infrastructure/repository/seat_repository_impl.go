package repository

import (
	"context"

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

func (r *seatRepository) Create(ctx context.Context, seat *entity.Seat) error {
	return r.db.WithContext(ctx).Create(seat).Error
}

func (r *seatRepository) FindByIDForUpdate(ctx context.Context, id uint) (*entity.Seat, error) {
	var m model.SeatModel
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}

	return &entity.Seat{
		ID:         m.ID,
		SeatNumber: m.SeatNumber,
		Status:     m.Status,
		Version:    m.Version,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}, nil
}

func (r *seatRepository) Update(ctx context.Context, seat *entity.Seat) error {
	return r.db.WithContext(ctx).Model(&model.SeatModel{}).Where("id=?", seat.ID).Update("status", seat.Status).Error
}
