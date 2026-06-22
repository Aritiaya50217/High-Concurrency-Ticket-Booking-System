package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	domainEvent "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
)

type BookingExpirationWorker struct {
	bookingRepo repository.BookingRepository
	outboxRepo  repository.OutboxRepository
}

func NewBookingExpirationWorker(bookingRepo repository.BookingRepository, outboxRepo repository.OutboxRepository) *BookingExpirationWorker {
	return &BookingExpirationWorker{bookingRepo: bookingRepo, outboxRepo: outboxRepo}
}

func (w *BookingExpirationWorker) Start() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		ctx := context.Background()

		bookings, err := w.bookingRepo.FindExpiredBookings(ctx)
		if err != nil {
			log.Println("find expired booking error : ", err)
			continue
		}

		for _, booking := range bookings {
			log.Println("booking expired : ", booking.ID)
			booking.Expire()

			if err := w.bookingRepo.UpdateStatus(ctx, booking.ID, string(booking.Status)); err != nil {
				log.Println("update booking fail : ", err)
				continue
			}

			cancelledEvent := domainEvent.NewBookingCancelled(booking.ID, booking.EventID, booking.SeatID)
			payloadBytes, err := json.Marshal(cancelledEvent)
			if err != nil {
				log.Println("marshal booking cancelled fail:", err)
				continue
			}

			outbox := &entity.OutboxEvent{
				EventType:   cancelledEvent.EventName(),
				Payload:     string(payloadBytes),
				Status:      entity.OutboxPending,
				CreatedAt:   time.Now(),
				ProcessedAt: time.Now(),
			}
			
			if err := w.outboxRepo.Create(ctx, outbox); err != nil {
				log.Println("create outbox fail:", err)
				continue
			}
			log.Printf("booking.cancelled created: booking=%d seat=%d", booking.ID, booking.SeatID)
		}
	}
}
