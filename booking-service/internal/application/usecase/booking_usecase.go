package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
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
	// call ticket-service
	token, ok := ctx.Value("token").(string)
	if !ok {
		return nil, errors.New("missing token")
	}
	fmt.Println("STEP 1 reserve seat")

	if err := u.eventClient.ReserveSeat(ctx, token, eventID, seatID); err != nil {
		return nil, err
	}

	fmt.Println("STEP 2 begin transaction")

	var result *aggregate.Booking

	err := u.bookingRepo.WithTransaction(ctx, func(repo repository.TxRepository) error {
		fmt.Println("STEP 3 create booking aggregate")

		booking := aggregate.NewBooking(userID, eventID, seatID)
		if booking == nil {
			return errors.New("booking is nil")
		}

		fmt.Println("booking:", booking)

		if err := repo.Booking().Create(ctx, booking); err != nil {
			fmt.Println("create booking fail:", err)
			return err
		}

		// domain event
		fmt.Println("STEP 4 process events")

		for _, e := range booking.Events() {
			fmt.Printf("event=%+v\n", e)

			if e == nil {
				continue
			}

			payload, err := json.Marshal(e)
			if err != nil {
				return err
			}
			fmt.Println("payload=", string(payload))

			outbox := &entity.OutboxEvent{
				EventType: e.EventName(),
				Payload:   string(payload),
				Status:    entity.OutboxPending,
				CreatedAt: time.Now(),
			}
			fmt.Printf("outbox=%+v\n", outbox)

			if err := repo.Outbox().Create(ctx, outbox); err != nil {
				fmt.Println("outbox fail:", err)
				return err
			}
		}

		// กัน publish ซ้ำ
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

// func (u *BookingUsecase) Update(id uint, status string) (*aggregate.Booking, error) {
// 	bookingID, err := u.bookingRepo.FindBookingByID(id)
// 	if err != nil {
// 		return nil, errors.New("get booking by ID not found")
// 	}

// 	booking := &aggregate.Booking{
// 		ID:     bookingID.ID,
// 		UserID: bookingID.UserID,
// 		// Status: status,
// 	}

// 	if err := u.bookingRepo.UpdateStatus(booking.ID, ""); err != nil {
// 		return nil, errors.New("error : update status ")
// 	}

// 	return booking, nil
// }

// func (u *BookingUsecase) Delete(id uint) error {
// 	bookingID, err := u.bookingRepo.FindBookingByID(id)
// 	if err != nil {
// 		return errors.New("get booking by ID not found")
// 	}

// 	if err := u.bookingRepo.Delete(bookingID.ID); err != nil {
// 		return errors.New("error : delete booking ")
// 	}
// 	return nil
// }

// func (u *BookingUsecase) FindByID(id uint) (*aggregate.Booking, error) {
// 	bookingID, err := u.bookingRepo.FindBookingByID(id)
// 	if err != nil {
// 		return nil, errors.New("get booking by ID not found")
// 	}

// 	return bookingID, nil
// }

// func (u *BookingUsecase) Search(status string, offset, limit int) ([]aggregate.Booking, int64, error) {
// 	return u.bookingRepo.Search(status, offset, limit)
// }
