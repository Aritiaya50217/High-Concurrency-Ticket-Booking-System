package repository

import (
	"context"
	"encoding/json"

	domainRepo "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/kafka"
)

type eventProducerRepository struct {
	producer *kafka.Producer
}

func NewEventProducerRepository(producer *kafka.Producer) domainRepo.EventProducerRepository {
	return &eventProducerRepository{producer: producer}
}

func (r *eventProducerRepository) Publish(ctx context.Context, event any) error {

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return r.producer.Publish(ctx, data)
}
