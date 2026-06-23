package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/dto"
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

	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("outbox worker stopped")
			return

		case <-ticker.C:
			events, err := w.repo.FindPending(context.Background())
			if err != nil {
				log.Println("find pending fail: ", err)
				continue
			}

			for _, event := range events {

				// validate payload
				if !json.Valid([]byte(event.Payload)) {
					log.Println("invalid playload :", event.Payload)
					continue
				}

				envelope := dto.EventEnvelope{
					Event: event.EventType,
					Data:  json.RawMessage([]byte(event.Payload)),
				}

				// publish
				err := w.producer.Publish(ctx, w.topic, envelope)
				if err != nil {
					log.Printf("publish failed topic=%s payload=%s err=%v", w.topic, event.Payload, err)
					continue
				}

				event.Status = entity.OutboxSent
				event.ProcessedAt = time.Now()

				if err := w.repo.Update(ctx, &event); err != nil {
					log.Println("update outbox fail: ", err)
				}

				log.Println("published: ", event.EventType)
			}
		}
	}

}
