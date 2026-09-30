package grpc_test

import (
	"context"
	"testing"

	userpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/user/v1"
	grpcserver "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/grpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUserServer_GetUser_InvalidUserID(t *testing.T) {
	server := grpcserver.NewUserServer(nil)

	_, err := server.GetUser(context.Background(), &userpb.GetUserRequest{UserId: 0})

	require.Equal(t, codes.InvalidArgument, status.Code(err))
}
