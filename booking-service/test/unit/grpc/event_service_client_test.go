package grpc_test

import (
	"context"
	"net"
	"testing"

	grpcclient "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/grpc"
	eventpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/event/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockEventServer struct {
	eventpb.UnimplementedEventServiceServer
}

type errorEventServer struct {
	eventpb.UnimplementedEventServiceServer
}

func (s *mockEventServer) ReserveSeat(ctx context.Context, req *eventpb.ReserveSeatRequest) (*eventpb.ReserveSeatResponse, error) {
	return &eventpb.ReserveSeatResponse{Success: true}, nil
}

func (s *errorEventServer) ReserveSeat(ctx context.Context, req *eventpb.ReserveSeatRequest) (*eventpb.ReserveSeatResponse, error) {
	return nil, status.Error(codes.NotFound, "event or seat not found")
}

func TestEventServiceClient_ReserveSeatSuccess(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	server := grpc.NewServer()
	eventpb.RegisterEventServiceServer(server, &mockEventServer{})

	go server.Serve(listener)
	defer server.Stop()

	addr := listener.Addr().String()

	client, err := grpcclient.NewEventServiceClient(addr)
	require.NoError(t, err)
	defer client.Close()

	eventID := uint(1)
	seatID := uint(10)
	userID := uint(100)

	err = client.ReserveSeat(context.Background(), eventID, seatID, userID)
	require.NoError(t, err)
}

func TestEventServiceClient_ReserveSeatError(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	server := grpc.NewServer()
	eventpb.RegisterEventServiceServer(server, &errorEventServer{})

	go server.Serve(listener)
	defer server.Stop()

	addr := listener.Addr().String()

	client, err := grpcclient.NewEventServiceClient(addr)
	require.NoError(t, err)
	defer client.Close()

	eventID := uint(1)
	seatID := uint(10)
	userID := uint(100)

	err = client.ReserveSeat(context.Background(), eventID, seatID, userID)
	require.Error(t, err)
	require.ErrorContains(t, err, "reserve seat via event service")
}
