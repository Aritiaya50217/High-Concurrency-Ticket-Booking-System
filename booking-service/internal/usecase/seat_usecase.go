package usecase

import (
	"context"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/repository"
)

type SeatUsecase struct {
	seatRepo repository.SeatRepository
}

func NewSeatUsecase(seatRepo repository.SeatRepository) *SeatUsecase {
	return &SeatUsecase{seatRepo: seatRepo}
}

func (u *SeatUsecase) Create(ctx context.Context, seat *entity.Seat) error {
	return u.seatRepo.Create(ctx, seat)
}
