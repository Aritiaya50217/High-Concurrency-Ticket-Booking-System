package worker

import (
	"context"
	"log"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
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

	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()

	for range ticker.C {
		events, err := w.repo.FindPending(context.Background())
		if err != nil {
			log.Println("find pending fail: ", err)
			continue
		}

		for _, event := range events {
			err := w.producer.Publish(context.Background(), event.EventType, event)
			if err != nil {
				log.Println("kafka publish fails: ", err)
				continue
			}
			now := time.Now()

			event.Status = entity.OutboxSent
			event.ProcessedAt = now

			if err := w.repo.Update(context.Background(), &event); err != nil {
				log.Println("update outbox fail: ", err)
			}

			log.Println("published: ", event.EventType)
		}
	}
}

// func (w *OutboxWorker) process(ctx context.Context) {
// 	events, err := w.repo.FindPending(ctx)
// 	if err != nil {
// 		log.Println("fetch outbox error : ", err)
// 		return
// 	}

// 	for _, event := range events {
// 		if err := w.producer.Publish(ctx, w.topic, json.RawMessage(event.Payload)); err != nil {
// 			log.Println("kafka publish fail : ", err)
// 			continue
// 		}

// 		if err := w.repo.MarkProcessed(ctx, event.ID); err != nil {
// 			log.Println("mark processed fail : ", err)
// 		}

// 	}
// }
