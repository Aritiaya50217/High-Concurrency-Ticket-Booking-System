package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
)

type BookingUsecase struct {
	bookingRepo repository.BookingRepository
	producer    *kafka.Producer
	topic       string
}

func NewBookingUsecase(bookingRepo repository.BookingRepository, producer *kafka.Producer, topic string) *BookingUsecase {
	return &BookingUsecase{bookingRepo: bookingRepo, producer: producer, topic: topic}
}

func (u *BookingUsecase) Create(ctx context.Context, userID, eventID, seatID uint) (*aggregate.Booking, error) {
	var result *aggregate.Booking

	err := u.bookingRepo.WithTransaction(ctx, func(repo repository.TxRepository) error {
		// business event
		booking := aggregate.NewBooking(userID, eventID, seatID)
		if err := repo.Booking().Create(ctx, booking); err != nil {
			return err
		}

		// domain event
		for _, e := range booking.Events() {
			if e == nil {
				continue
			}

			payload, err := json.Marshal(e)
			if err != nil {
				return nil
			}

			outbox := &entity.OutboxEvent{
				EventType: e.EventName(),
				Payload:   string(payload),
				Status:    "PENDING",
			}
			if err := repo.Outbox().Create(ctx, outbox); err != nil {
				return err
			}
		}

		result = booking
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (u *BookingUsecase) Update(id uint, status string) (*aggregate.Booking, error) {
	bookingID, err := u.bookingRepo.FindBookingByID(id)
	if err != nil {
		return nil, errors.New("get booking by ID not found")
	}

	booking := &aggregate.Booking{
		ID:     bookingID.ID,
		UserID: bookingID.UserID,
		// Status: status,
	}

	if err := u.bookingRepo.UpdateStatus(booking.ID, ""); err != nil {
		return nil, errors.New("error : update status ")
	}

	return booking, nil
}

func (u *BookingUsecase) Delete(id uint) error {
	bookingID, err := u.bookingRepo.FindBookingByID(id)
	if err != nil {
		return errors.New("get booking by ID not found")
	}

	if err := u.bookingRepo.Delete(bookingID.ID); err != nil {
		return errors.New("error : delete booking ")
	}
	return nil
}

func (u *BookingUsecase) FindByID(id uint) (*aggregate.Booking, error) {
	bookingID, err := u.bookingRepo.FindBookingByID(id)
	if err != nil {
		return nil, errors.New("get booking by ID not found")
	}

	return bookingID, nil
}

func (u *BookingUsecase) Search(status string, offset, limit int) ([]aggregate.Booking, int64, error) {
	return u.bookingRepo.Search(status, offset, limit)
}
