package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
)

type BookingUsecase struct {
	bookingRepo repository.BookingRepository
	seatRepo    repository.SeatRepository
	producer    *kafka.Producer
	topic       string
}

func NewBookingUsecase(bookingRepo repository.BookingRepository, seatRepo repository.SeatRepository, producer *kafka.Producer, topic string) *BookingUsecase {
	return &BookingUsecase{bookingRepo: bookingRepo, seatRepo: seatRepo, producer: producer, topic: topic}
}

func (u *BookingUsecase) Create(ctx context.Context, userID, eventID, seatID uint) (*entity.Booking, error) {
	var result *entity.Booking
	err := u.bookingRepo.WithTransaction(ctx, func(repo repository.TxRepository) error {
		seat, err := repo.Seat().FindByID(ctx, seatID)
		if err != nil {
			return err
		}

		if seat.Status != entity.SeatAvailable {
			return errors.New("seat already taken")
		}

		booking := &entity.Booking{
			UserID:  userID,
			EventID: eventID,
			SeatID:  seatID,
			Status:  entity.SeatBooked,
		}

		if err := repo.Booking().Create(ctx, booking); err != nil {
			return err
		}

		if err := repo.Seat().MarkBooked(ctx, seatID); err != nil {
			return err
		}

		// publish kafka
		payload, _ := json.Marshal(map[string]interface{}{
			"event":      "BOOKING_CREATED",
			"booking_id": result.ID,
			"user_id":    result.UserID,
			"event_id":   result.EventID,
			"seat_id":    result.SeatID,
		})

		outbox := &entity.OutboxEvent{
			EventType: "BOOKING_CREATED",
			Payload:   string(payload),
			Status:    "PENDING",
		}

		if err := repo.Outbox().Create(ctx, outbox); err != nil {
			return err
		}

		result = booking
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (u *BookingUsecase) Update(id uint, status string) (*entity.Booking, error) {
	bookingID, err := u.bookingRepo.FindBookingByID(id)
	if err != nil {
		return nil, errors.New("get booking by ID not found")
	}

	booking := &entity.Booking{
		ID:     bookingID.ID,
		UserID: bookingID.UserID,
		Status: status,
	}

	if err := u.bookingRepo.UpdateStatus(booking.ID, booking.Status); err != nil {
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

func (u *BookingUsecase) FindByID(id uint) (*entity.Booking, error) {
	bookingID, err := u.bookingRepo.FindBookingByID(id)
	if err != nil {
		return nil, errors.New("get booking by ID not found")
	}

	return bookingID, nil
}

func (u *BookingUsecase) Search(status string, offset, limit int) ([]entity.Booking, int64, error) {
	return u.bookingRepo.Search(status, offset, limit)
}
