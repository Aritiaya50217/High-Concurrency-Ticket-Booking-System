package grpc

import (
	"context"

	userpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	client userpb.UserServiceClient
	conn   *grpc.ClientConn
}

func NewUserServiceClient(addr string) (*UserServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(conn)

	return &UserServiceClient{client: client, conn: conn}, nil
}

func (c *UserServiceClient) GetUser(ctx context.Context, userID uint64) (bool, error) {
	resp, err := c.client.GetUser(ctx, &userpb.GetUserRequest{
		UserId: userID,
	})
	if err != nil {
		return false, err
	}

	return resp.Exists, nil
}

func (c *UserServiceClient) Close() error {
	return c.conn.Close()
}
