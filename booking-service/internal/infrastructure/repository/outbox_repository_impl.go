package repository

import (
	"context"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

type outboxRepository struct {
	db *gorm.DB
}

func NewOutboxRepository(db *gorm.DB) repository.OutboxRepository {
	return &outboxRepository{db: db}
}

func (r *outboxRepository) Create(ctx context.Context, outbox *entity.OutboxEvent) error {
	return r.db.WithContext(ctx).Create(outbox).Error
}

func (r *outboxRepository) FindPending(ctx context.Context, limit int) ([]entity.OutboxEvent, error) {
	var events []model.OutboxEventModel
	err := r.db.WithContext(ctx).Where("status = ? ", "PENDING").Limit(limit).Find(&events).Error
	results := []entity.OutboxEvent{}
	for _, e := range events {
		results = append(results, entity.OutboxEvent{
			ID:          e.ID,
			EventType:   e.EventType,
			Payload:     e.Payload,
			Status:      e.Status,
			CreatedAt:   e.CreatedAt,
			ProcessedAt: e.ProcessedAt,
		})
	}

	return results, err
}

func (r *outboxRepository) MarkProcessed(ctx context.Context, id string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.OutboxEventModel{}).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status":       "PROCESSED",
		"processed_at": now,
	}).Error
}
