package grpc

import (
	"context"
	"log"

	eventpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/event/v1"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/application/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EventServer struct {
	eventpb.UnimplementedEventServiceServer
	eventUsecase *usecase.EventUsecase
}

func NewEventServer(eventUsecase *usecase.EventUsecase) *EventServer {
	return &EventServer{
		eventUsecase: eventUsecase,
	}
}

func (s *EventServer) ReserveSeat(ctx context.Context, req *eventpb.ReserveSeatRequest) (*eventpb.ReserveSeatResponse, error) {

	if req.EventId == 0 {
		return nil, status.Error(codes.InvalidArgument, "event_id is required")
	}

	if req.SeatId == 0 {
		return nil, status.Error(codes.InvalidArgument, "seat_id is required")
	}

	if req.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}
	
	eventID := uint(req.EventId)
	seatID := uint(req.SeatId)
	userID := uint(req.UserId)

	log.Printf(
		"gRPC ReserveSeat called: event_id=%d seat_id=%d user_id=%d",
		req.EventId,
		req.SeatId,
		req.UserId,
	)

	if err := s.eventUsecase.ReserveSeat(ctx, eventID, seatID, userID); err != nil {
		return nil, err
	}

	return &eventpb.ReserveSeatResponse{Success: true}, nil
}
