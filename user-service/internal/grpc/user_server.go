package grpc

import (
	"context"

	userpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/user/v1"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/application/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	userUsecase *usecase.UserUsecase
}

func NewUserServer(userUsecase *usecase.UserUsecase) *UserServer {
	return &UserServer{userUsecase: userUsecase}
}

func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	if req.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	user, err := s.userUsecase.Profile(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &userpb.GetUserResponse{Exists: false}, nil
	}

	return &userpb.GetUserResponse{Exists: true}, nil
}
