package repository

import (
	"context"
	"errors"

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
			Version:    seat.Version,
		})
	}
	if err := r.db.WithContext(ctx).Create(&eventModel).Error; err != nil {
		return err
	}

	event.ID = eventModel.ID

	return nil

}

func (r *eventRepository) Transaction(ctx context.Context, fn func(repo domainRepo.EventRepository) error) error {
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
			Version:    seat.Version,
		})
	}

	return event, nil
}

func (r *eventRepository) Update(ctx context.Context, event *aggregate.Event) error {
	for _, seat := range event.Seats {
		result := r.db.WithContext(ctx).Model(&model.SeatModel{}).Where("id=? and version=?", seat.ID, seat.Version).
			Updates(map[string]interface{}{
				"status":  seat.Status,
				"version": seat.Version + 1},
			)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("seat updated by another transaction")
		}

		seat.Version++
	}
	return nil
}

func (r *eventRepository) FindByIDForUpdate(ctx context.Context, eventID uint) (*aggregate.Event, error) {
	var eventModel model.EventModel

	err := r.db.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
	}).Preload("Seats").First(&eventModel, eventID).Error

	if err != nil {
		return nil, err
	}

	event := &aggregate.Event{
		ID:   eventModel.ID,
		Name: eventModel.Name,
	}

	for _, seat := range event.Seats {
		event.Seats = append(event.Seats, &entity.Seat{
			ID:         seat.ID,
			EventID:    seat.EventID,
			SeatNumber: seat.SeatNumber,
			Status:     valueobject.SeatStatus(seat.Status),
			Version:    seat.Version,
		})
	}

	return event, nil
}
