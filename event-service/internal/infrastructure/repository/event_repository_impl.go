package repository

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/entity"
	domainRepo "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/valueobject"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) domainRepo.EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) Create(ctx context.Context, event *aggregate.Event) error {
	eventModel := model.EventModel{Name: event.Name}

	for _, seat := range event.Seats {
		eventModel.Seats = append(eventModel.Seats, model.SeatModel{
			SeatNumber: seat.SeatNumber,
			Status:     string(seat.Status),
		})
	}

	return r.db.WithContext(ctx).Create(&eventModel).Error
}

func (r *eventRepository) WithTransaction(ctx context.Context, fn func(repo domainRepo.EventRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &eventRepository{db: tx}
		return fn(txRepo)
	})
}

func (r *eventRepository) FindByID(ctx context.Context, id uint) (*aggregate.Event, error) {
	var eventModel model.EventModel

	if err := r.db.WithContext(ctx).Preload("Seats").First(&eventModel, id).Error; err != nil {
		return nil, err
	}

	event := &aggregate.Event{
		ID:   eventModel.ID,
		Name: eventModel.Name,
	}

	for _, seat := range eventModel.Seats {
		event.Seats = append(event.Seats, &entity.Seat{
			ID:         seat.ID,
			EventID:    seat.EventID,
			SeatNumber: seat.SeatNumber,
			Status:     valueobject.SeatStatus(seat.Status),
		})
	}

	return event, nil
}

func (r *eventRepository) Update(ctx context.Context, event *aggregate.Event) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, seat := range event.Seats {
			if err := tx.Model(&model.SeatModel{}).Where("id=?", seat.ID).Update("status", string(seat.Status)).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *eventRepository) FindByIDForUpdate(ctx context.Context, eventID uint) (*aggregate.Event, error) {
	var event aggregate.Event
	err := r.db.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
	}).Preload("Seats").First(&event, eventID).Error

	return &event, err
}
