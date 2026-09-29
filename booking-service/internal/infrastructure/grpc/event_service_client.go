package grpc

import (
	"context"

	eventpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/event/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EventServiceClient struct {
	client eventpb.EventServiceClient
	conn   *grpc.ClientConn
}

func NewEventServiceClient(addr string) (*EventServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := eventpb.NewEventServiceClient(conn)

	return &EventServiceClient{
		client: client,
		conn:   conn,
	}, nil

}

func (c *EventServiceClient) ReserveSeat(ctx context.Context, eventID, seatID, userID uint) error {
	_, err := c.client.ReserveSeat(ctx, &eventpb.ReserveSeatRequest{
		EventId: uint64(eventID),
		SeatId:  uint64(seatID),
		UserId:  uint64(userID),
	})
	return err
}

func (c *EventServiceClient) Close() error {
	return c.conn.Close()
}
