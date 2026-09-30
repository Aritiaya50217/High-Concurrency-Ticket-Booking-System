package grpc_test

import (
	"context"
	"net"
	"testing"

	grpcclient "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/grpc"
	userpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/user/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockUserServer struct {
	userpb.UnimplementedUserServiceServer
}

type errorUserServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *mockUserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{Exists: true}, nil
}

func (s *errorUserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return nil, status.Error(codes.NotFound, "user not found")
}

func TestUserServiceClient_GetUserSuccess(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	// gRPC server infrastructure
	server := grpc.NewServer()

	// เอา mockUserServer ไปลงทะเบียนกับ gRPC server ตัวนี้
	userpb.RegisterUserServiceServer(server, &mockUserServer{})

	go server.Serve(listener)
	defer server.Stop()

	addr := listener.Addr().String()

	// Client ที่อยู่ใน Booking Service
	client, err := grpcclient.NewUserServiceClient(addr)
	require.NoError(t, err)
	defer client.Close()

	userID := uint64(1)

	exists, err := client.GetUser(context.Background(), userID)

	require.NoError(t, err)
	require.NotNil(t, exists)
	require.True(t, exists)

}

func TestUserServiceClient_GetUserError(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, &errorUserServer{})

	go server.Serve(listener)
	defer server.Stop()

	addr := listener.Addr().String()

	client, err := grpcclient.NewUserServiceClient(addr)
	require.NoError(t, err)
	defer client.Close()

	userID := uint64(1)

	exists, err := client.GetUser(context.Background(), userID)

	require.Error(t, err)
	require.False(t, exists)
	require.ErrorContains(t, err, "get user via user service")
}
