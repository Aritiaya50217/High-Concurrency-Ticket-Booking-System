package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	domainEvent "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/external/eventservice"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
)

type BookingUsecase struct {
	bookingRepo repository.BookingRepository
	producer    *kafka.Producer
	topic       string
	eventClient eventservice.SeatServiceClient
}

func NewBookingUsecase(bookingRepo repository.BookingRepository, producer *kafka.Producer, topic string, eventClient eventservice.SeatServiceClient) *BookingUsecase {
	return &BookingUsecase{bookingRepo: bookingRepo, producer: producer, topic: topic, eventClient: eventClient}
}

func (u *BookingUsecase) Create(ctx context.Context, userID, eventID, seatID uint) (*aggregate.Booking, error) {

	fmt.Println("STEP 1 create booking")

	var result *aggregate.Booking

	err := u.bookingRepo.WithTransaction(ctx, func(repo repository.TxRepository) error {

		booking := aggregate.NewBooking(userID, eventID, seatID)
		if booking == nil {
			return errors.New("booking is nil")
		}

		if err := repo.Booking().Create(ctx, booking); err != nil {
			fmt.Println("Booking Create error : ", err)
			return err
		}

		log.Println("booking id =", booking.ID)
		
		log.Printf("domain events: %+v\n", booking.Events())

		fmt.Println("STEP 2 store outbox events")

		for _, e := range booking.Events() {

			if e == nil {
				log.Println("skip nil event")
				continue
			}

			payloadBytes, err := json.Marshal(e)
			if err != nil {
				return fmt.Errorf("marshal payload failed: %w", err)
			}

			if len(payloadBytes) == 0 || string(payloadBytes) == "null" || string(payloadBytes) == "{}" {
				return errors.New("invalid domain event payload")
			}

			outbox := &entity.OutboxEvent{
				EventType: e.EventName(),
				Payload:   string(payloadBytes), // keep simple, no new struct
				Status:    entity.OutboxPending,
				CreatedAt: time.Now(),
			}

			log.Printf("outbox create: %+v\n", outbox)

			if err := repo.Outbox().Create(ctx, outbox); err != nil {
				return fmt.Errorf("outbox create failed: %w", err)
			}
		}

		booking.ClearEvents()
		result = booking
		return nil
	})

	if err != nil {
		fmt.Println("transaction fail:", err)
		return nil, err
	}

	return result, nil
}

func (u *BookingUsecase) HandleSeatReserved(ctx context.Context, event domainEvent.SeatReserved) error {
	log.Println("received seat.reserved event:", event)

	// find booking
	booking, err := u.bookingRepo.FindByEventAndSeat(ctx, event.EventID, event.SeatID)
	if err != nil {
		return err
	}

	if booking == nil {
		return errors.New("booking not found")
	}

	if booking.Status == valueobject.BookingConfirmed {
		log.Println("booking already confirmed, skip")
		return nil
	}

	// update state
	booking.Confirm()

	// save
	if err := u.bookingRepo.UpdateStatus(booking.ID, string(booking.Status)); err != nil {
		return err
	}

	log.Println("booking confirmed : ", booking.ID)

	return nil
}
