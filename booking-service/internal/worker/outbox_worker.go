package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
)

type OutboxWorker struct {
	repo     repository.OutboxRepository
	producer *kafka.Producer
	topic    string
}

func NewOutboxWorker(repo repository.OutboxRepository, producer *kafka.Producer, topic string) *OutboxWorker {
	return &OutboxWorker{
		repo: repo, producer: producer, topic: topic,
	}
}

func (w *OutboxWorker) Start(ctx context.Context) {
	log.Println("outbox worker started")

	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			w.process(ctx)

		case <-ctx.Done():
			return
		}
	}
}

func (w *OutboxWorker) process(ctx context.Context) {
	events, err := w.repo.FindPending(ctx, 10)
	if err != nil {
		log.Println("fetch outbox error : ", err)
		return
	}

	for _, event := range events {
		if err := w.producer.Publish(ctx, w.topic, json.RawMessage(event.Payload)); err != nil {
			log.Println("kafka publish fail : ", err)
			continue
		}

		if err := w.repo.MarkProcessed(ctx, event.ID); err != nil {
			log.Println("mark processed fail : ", err)
		}

	}
}
