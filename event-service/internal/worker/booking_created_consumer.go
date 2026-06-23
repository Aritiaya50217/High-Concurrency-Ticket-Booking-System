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

		var envelope dto.EventEnvelope

		if err := json.Unmarshal(data, &envelope); err != nil {
			log.Println("unmarshal envelope fail:", err)
			return err
		}

		switch envelope.Event {
		case "booking.created":
			var event domainEvent.BookingCreated

			if err := json.Unmarshal(envelope.Data, &event); err != nil {
				log.Println("unmarshal booking fail:", err)
				return err
			}

			log.Printf("booking received : %+v", event)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			defer cancel()

			log.Println("calling usecase HandleBookingCreated")

			return b.usecase.HandleBookingCreated(ctx, event)

		case "booking.cancelled":
			var event domainEvent.BookingCancelled

			if err := json.Unmarshal(envelope.Data, &event); err != nil {
				log.Println("unmarshal booking.cancelled fail:", err)
				return err
			}

			return b.usecase.HandleBookingCancelled(context.Background(), event)

		default:
			return nil
		}
	})
}
