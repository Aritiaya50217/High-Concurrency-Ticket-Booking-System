package repository

import (
	"context"

	domainRepo "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/kafka"
)

type eventProducerRepository struct {
	producer *kafka.Producer
}

func NewEventProducerRepository(producer *kafka.Producer) domainRepo.EventProducerRepository {
	return &eventProducerRepository{producer: producer}
}

func (r *eventProducerRepository) Publish(ctx context.Context, topic string, message []byte) error {
	return r.producer.Publish(ctx, topic, message)
}
