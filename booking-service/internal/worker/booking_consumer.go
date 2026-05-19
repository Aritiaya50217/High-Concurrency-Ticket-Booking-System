package worker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/dto"
)

type BookingConsumer struct {
	consumer *kafka.Consumer
}

func NewBookingConsumer(c *kafka.Consumer) *BookingConsumer {
	return &BookingConsumer{consumer: c}
}

func (w *BookingConsumer) Start(ctx context.Context) {
	log.Println("booking consumer started ...")
	
	w.consumer.Start(ctx, func(msg []byte) {
		var event dto.BookingCreatedEvent
		if err := json.Unmarshal(msg, &event); err != nil {
			log.Println("invalid event format : ", err)
			return
		}

		switch event.Event {
		case "booking.created":
			w.handleBookingCreated(ctx, event)
		default:
			log.Println("unknown event : ", event.Event)
		}
	})
}

func (w *BookingConsumer) handleBookingCreated(ctx context.Context, event dto.BookingCreatedEvent) {
	log.Println("booking.created received: ", "booking_id=", event.BookingID,
		"user_id=", event.UserID,
		// "event_id=", event.EventID,
	)

	log.Println("processed booking event:", event.BookingID)
}
