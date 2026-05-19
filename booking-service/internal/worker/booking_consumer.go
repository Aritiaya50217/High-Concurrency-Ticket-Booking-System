package worker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/application/usecase"
	domainEvent "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/dto"
)

type BookingConsumer struct {
	consumer *kafka.Consumer
	usecase  *usecase.BookingUsecase
}

func NewBookingConsumer(c *kafka.Consumer, usecase *usecase.BookingUsecase) *BookingConsumer {
	return &BookingConsumer{consumer: c, usecase: usecase}
}

func (w *BookingConsumer) Start(ctx context.Context) {

	log.Println("booking consumer started ...")

	w.consumer.Start(ctx, func(msg []byte) {

		var envelope dto.EventEnvelope

		if err := json.Unmarshal(msg, &envelope); err != nil {
			log.Println("invalid envelope:", err)
			return
		}

		switch envelope.Event {

		case "booking.created":
			var event dto.BookingCreatedEvent
			if err := json.Unmarshal(envelope.Data, &event); err != nil {
				log.Println("invalid booking.created:", err)
				return
			}
			w.handleBookingCreated(ctx, event)

		case "seat.reserved":
			var event dto.SeatReservedEvent
			if err := json.Unmarshal(envelope.Data, &event); err != nil {
				log.Println("invalid seat.reserved:", err)
				return
			}
			w.handleSeatReserved(ctx, event)

		default:
			log.Println("unknown event:", envelope.Event)
		}
	})
}

func (w *BookingConsumer) handleBookingCreated(ctx context.Context, event dto.BookingCreatedEvent) {
	log.Println("booking.created received: ", "booking_id=", event.BookingID, "user_id=", event.UserID)
	log.Println("processed booking event:", event.BookingID)
}

func (w *BookingConsumer) handleSeatReserved(ctx context.Context, event dto.SeatReservedEvent) {
	log.Println("seat.reserved received:", "booking_id=", event.BookingID, "seat_id=", event.SeatID)

	seat := domainEvent.SeatReserved{
		BookingID: event.BookingID,
		EventID:   event.EventID,
		SeatID:    event.SeatID,
	}

	if err := w.usecase.HandleSeatReserved(ctx, seat); err != nil {
		log.Println("update booking failed:", err)
		return
	}

	log.Println("booking updated to CONFIRMED:", event.BookingID)
}
