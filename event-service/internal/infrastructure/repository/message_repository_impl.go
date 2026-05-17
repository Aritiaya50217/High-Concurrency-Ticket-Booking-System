package repository

import (
	"context"

	domainRepo "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/kafka"
)

type messageRepository struct {
	producer *kafka.Producer
}

func NewMessageRepository(producer *kafka.Producer) domainRepo.MessageRepository {
	return &messageRepository{producer: producer}
}

func (m *messageRepository) Publish(ctx context.Context, event any) error {
	return m.producer.Publish(ctx, event)
}
