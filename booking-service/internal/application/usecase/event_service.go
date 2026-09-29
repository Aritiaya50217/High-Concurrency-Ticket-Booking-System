package usecase

import "context"

type EventService interface {
	ReserveSeat(ctx context.Context, eventID, seatID, userID uint) error
}
