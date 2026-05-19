package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
	"gorm.io/gorm"
)

type inboxRepository struct {
	db *gorm.DB
}

func NewInboxRepository(db *gorm.DB) repository.InboxRepository {
	return &inboxRepository{db: db}
}

func (r *inboxRepository) IsProcessed(ctx context.Context, eventID string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Table("processed_events").Where("event_id = ?", eventID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *inboxRepository) MarkProcessed(ctx context.Context, eventID string, eventType string) error {
	return r.db.WithContext(ctx).Table("processed_events").Create(map[string]interface{}{
		"event_id":   eventID,
		"event_type": eventType,
	}).Error
}
