package grpc_test

import (
	"context"
	"testing"

	eventpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/event/v1"
	grpcserver "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/grpc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestEventServer_ReserveSeat_InvalidEventID(t *testing.T) {
	server := grpcserver.NewEventServer(nil)

	_, err := server.ReserveSeat(context.Background(), &eventpb.ReserveSeatRequest{
		EventId: 0,
		SeatId:  1,
		UserId:  1,
	})

	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}


func TestEventServer_ReserveSeat_InvalidSeatID(t *testing.T) {
	server := grpcserver.NewEventServer(nil)

	_, err := server.ReserveSeat(context.Background(), &eventpb.ReserveSeatRequest{
		EventId: 1,
		SeatId:  0,
		UserId:  1,
	})

	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestEventServer_ReserveSeat_InvalidUserID(t *testing.T) {
	server := grpcserver.NewEventServer(nil)

	_, err := server.ReserveSeat(context.Background(), &eventpb.ReserveSeatRequest{
		EventId: 1,
		SeatId:  1,
		UserId:  0,
	})

	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}
