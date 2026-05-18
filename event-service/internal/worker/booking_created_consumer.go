package worker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/application/usecase"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/kafka"

	domainEvent "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/event"
)

type BookingCreatedConsumer struct {
	consumer *kafka.Consumer
	usecase  *usecase.EventUsecase
}

func NewBookingCreatedConsumer(consumer *kafka.Consumer, usecase *usecase.EventUsecase) *BookingCreatedConsumer {
	return &BookingCreatedConsumer{consumer: consumer, usecase: usecase}
}

func (b *BookingCreatedConsumer) Start() {
	b.consumer.Consume(context.Background(), func(data []byte) error {
		var event domainEvent.BookingCreated

		if err := json.Unmarshal(data, &event); err != nil {
			return err
		}
		log.Printf("booking received : %+v", event)

		return b.usecase.HandleBookingCreated(context.Background(), event)
	})
}
