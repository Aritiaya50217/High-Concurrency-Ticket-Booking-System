package grpc

import (
	"context"

	userpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/user"
	usecase "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/usecase"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	userUsecase *usecase.UserUsecase
}

func NewUserServer(userUsecase *usecase.UserUsecase) *UserServer {
	return &UserServer{userUsecase: userUsecase}
}

func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := s.userUsecase.Profile(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &userpb.GetUserResponse{Exists: false}, nil
	}

	return &userpb.GetUserResponse{Exists: true}, nil
}
