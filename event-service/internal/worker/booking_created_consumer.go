package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/dto"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/kafka"

	domainEvent "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/metrics"
)

type BookingCreatedConsumer struct {
	consumer  *kafka.Consumer
	usecase   *usecase.EventUsecase
	inboxRepo repository.InboxRepository
}

func NewBookingCreatedConsumer(consumer *kafka.Consumer, usecase *usecase.EventUsecase, inboxRepo repository.InboxRepository) *BookingCreatedConsumer {
	return &BookingCreatedConsumer{consumer: consumer, usecase: usecase, inboxRepo: inboxRepo}
}

func (b *BookingCreatedConsumer) Start() {
	b.consumer.Start(context.Background(), func(data []byte) error {
		log.Println("received message:", string(data))

		metrics.KafkaMessagesConsumedTotal.Inc()

		var envelope dto.EventEnvelope

		if err := json.Unmarshal(data, &envelope); err != nil {
			log.Println("unmarshal envelope fail:", err)

			metrics.KafkaConsumerErrorsTotal.WithLabelValues("unmarshal").Inc()

			return err
		}

		switch envelope.Event {
		case "booking.created":
			var event domainEvent.BookingCreated

			if err := json.Unmarshal(envelope.Data, &event); err != nil {
				log.Println("unmarshal booking fail:", err)

				metrics.KafkaConsumerErrorsTotal.WithLabelValues("unmarshal").Inc()

				return err
			}

			log.Printf("booking received : %+v", event)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			defer cancel()

			log.Println("calling usecase HandleBookingCreated")

			if err := b.usecase.HandleBookingCreated(ctx, event); err != nil {
				metrics.KafkaConsumerErrorsTotal.WithLabelValues("usecase").Inc()
				return err
			}

			metrics.KafkaMessagesProcessedTotal.Inc()
			return nil

		case "booking.cancelled":
			var event domainEvent.BookingCancelled

			if err := json.Unmarshal(envelope.Data, &event); err != nil {
				log.Println("unmarshal booking.cancelled fail:", err)

				metrics.KafkaConsumerErrorsTotal.WithLabelValues("unmarshal").Inc()

				return err
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			defer cancel()

			if err := b.usecase.HandleBookingCancelled(ctx, event); err != nil {
				metrics.KafkaConsumerErrorsTotal.WithLabelValues("usecase").Inc()
				return err
			}

			metrics.KafkaMessagesProcessedTotal.Inc()

			return nil

		default:
			return nil
		}
	})
}
