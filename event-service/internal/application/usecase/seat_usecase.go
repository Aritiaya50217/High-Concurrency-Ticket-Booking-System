package usecase

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/domain/repository"
)

type SeatUsecase struct {
	repo repository.EventRepository
}

func NewSeatUsecase(repo repository.EventRepository) *SeatUsecase {
	return &SeatUsecase{repo: repo}
}

func (u *SeatUsecase) Reserve(ctx context.Context, eventID, seatID uint) error {
	return u.repo.Transaction(ctx, func(repo repository.EventRepository) error {
		event, err := repo.FindByIDForUpdate(ctx, eventID)
		if err != nil {
			return err
		}

		err = event.ReserveSeat(seatID)
		if err != nil {
			return err
		}
		return repo.Update(ctx, event)
	})
}

func (u *SeatUsecase) Release(ctx context.Context, eventID, seatID uint) error {
	return u.repo.Transaction(ctx, func(repo repository.EventRepository) error {
		event, err := repo.FindByIDForUpdate(ctx, eventID)
		if err != nil {
			return err
		}

		err = event.ReleaseSeat(seatID)
		if err != nil {
			return err
		}

		return repo.Update(ctx, event)
	})
}
