package repository

import (
	"context"
	"fmt"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/database/model"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type outboxRepository struct {
	db *gorm.DB
}

func NewOutboxRepository(db *gorm.DB) repository.OutboxRepository {
	return &outboxRepository{db: db}
}

func (r *outboxRepository) Create(ctx context.Context, outbox *entity.OutboxEvent) error {
	outboxModel := model.OutboxEventModel{
		EventType: outbox.EventType,
		Payload: datatypes.JSON(
			[]byte(outbox.Payload),
		),
		Status:      outbox.Status,
		CreatedAt:   outbox.CreatedAt,
		ProcessedAt: outbox.ProcessedAt,
	}
	err := r.db.
		WithContext(ctx).
		Create(&outboxModel).
		Error

	fmt.Println("gorm err =", err)

	return err
}

func (r *outboxRepository) FindPending(ctx context.Context) ([]entity.OutboxEvent, error) {
	var events []model.OutboxEventModel
	err := r.db.WithContext(ctx).Where("status = ? ", "PENDING").Find(&events).Error
	results := []entity.OutboxEvent{}
	for _, e := range events {
		results = append(results, entity.OutboxEvent{
			ID:          e.ID,
			EventType:   e.EventType,
			Payload:     "",
			Status:      e.Status,
			CreatedAt:   e.CreatedAt,
			ProcessedAt: e.ProcessedAt,
		})
	}

	return results, err
}

func (r *outboxRepository) Update(ctx context.Context, outbox *entity.OutboxEvent) error {
	return r.db.WithContext(ctx).Model(&model.OutboxEventModel{}).Where("id=?", outbox.ID).
		Updates(map[string]interface{}{
			"status":       outbox.Status,
			"processed_at": outbox.ProcessedAt,
		}).Error
}
