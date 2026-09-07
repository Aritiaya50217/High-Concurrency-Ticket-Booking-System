package integration

import (
	"context"
	"net"
	"testing"

	userpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	userusecase "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/entity"
	usergrpc "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/grpc"
	usermocks "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/test/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func startTestGRPCServer(t *testing.T, repo *usermocks.MockUserRepository) string {
	t.Helper()

	userUsecase := userusecase.NewUserUsecase(repo, nil)

	server := grpc.NewServer()

	userServer := usergrpc.NewUserServer(userUsecase)

	userpb.RegisterUserServiceServer(server, userServer)

	listener, err := net.Listen("tcp", "127.0.0.1:0")

	require.NoError(t, err)

	go func() {
		_ = server.Serve(listener)
	}()

	t.Cleanup(func() {
		server.Stop()
		listener.Close()
	})

	return listener.Addr().String()

}

func createGRPCClient(t *testing.T, addr string) userpb.UserServiceClient {
	t.Helper()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	require.NoError(t, err)

	t.Cleanup(func() { conn.Close() })

	return userpb.NewUserServiceClient(conn)
}

func TestGetUserGRPCUserExsits(t *testing.T) {
	repo := new(usermocks.MockUserRepository)

	userID := uint(1)

	user := &entity.Users{
		ID: userID,
	}

	repo.On("Profile", mock.AnythingOfType("uint")).Return(user, nil)

	addr := startTestGRPCServer(t, repo)

	client := createGRPCClient(t, addr)

	ctx := context.Background()

	resp, err := client.GetUser(ctx, &userpb.GetUserRequest{UserId: 1})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.True(t, resp.Exists)

	repo.AssertExpectations(t)

}

func TestGetUserGRPCUserNotFound(t *testing.T) {
	repo := new(usermocks.MockUserRepository)

	repo.On("Profile", mock.AnythingOfType("uint")).Return(nil, nil)

	addr := startTestGRPCServer(t, repo)

	client := createGRPCClient(t, addr)

	ctx := context.Background()

	resp, err := client.GetUser(ctx, &userpb.GetUserRequest{UserId: 999})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.False(t, resp.Exists)

	repo.AssertExpectations(t)
}

func TestGetUserGRPCRepositoryError(t *testing.T) {
	repo := new(usermocks.MockUserRepository)

	repo.On("Profile", mock.AnythingOfType("uint")).Return(nil, context.DeadlineExceeded)

	addr := startTestGRPCServer(t, repo)

	client := createGRPCClient(t, addr)

	ctx := context.Background()

	resp, err := client.GetUser(ctx, &userpb.GetUserRequest{UserId: 1})

	require.Error(t, err)
	require.Nil(t, resp)

	repo.AssertExpectations(t)
}
