package grpc_test

import (
	"context"
	"net"
	"testing"
	"time"

	eventpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/event/v1"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/application/usecase"
	eventgrpc "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/grpc"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/database"
	infraRepo "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/gorm"
)

const bufSize = 1024 * 1024

func TestReserveSeatGRPCSuccess(t *testing.T) {
	cfg := &config.Config{}

	cfg.Database.Host = "localhost"
	cfg.Database.Port = 5433
	cfg.Database.User = "admin"
	cfg.Database.Password = "password123!"
	cfg.Database.Name = "event_db"
	cfg.Database.SSLMode = "disable"

	db, err := database.NewPostgresDB(cfg)
	require.NoError(t, err)

	seedEvent(t, db)

	producer := new(mocks.MockEventProducer)

	repo := infraRepo.NewEventRepository(db)

	u := usecase.NewEventUsecase(repo, nil, producer)

	server := eventgrpc.NewEventServer(u)

	listener := bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()

	eventpb.RegisterEventServiceServer(grpcServer, server)

	go func() {
		_ = grpcServer.Serve(listener)
	}()

	t.Cleanup(func() {
		grpcServer.Stop()
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		conn.Close()
	})

	client := eventpb.NewEventServiceClient(conn)
	producer.On("Publish", mock.Anything, "seat.reserved", mock.Anything).Return(nil)

	resp, err := client.ReserveSeat(ctx, &eventpb.ReserveSeatRequest{EventId: 1, SeatId: 1, UserId: 1})

	require.NoError(t, err)
	require.True(t, resp.Success)

	var seat struct {
		Status  string
		Version int
	}

	err = db.Table("seats").Where("id = ?", 1).First(&seat).Error

	require.NoError(t, err)

	require.Equal(t, "RESERVED", seat.Status)
	require.Equal(t, 1, seat.Version)

	producer.AssertExpectations(t)
}

func seedEvent(t *testing.T, db *gorm.DB) {
	t.Helper()

	err := db.Exec("DELETE FROM seats WHERE event_id = ?", 1).Error
	require.NoError(t, err)

	err = db.Exec("DELETE FROM events WHERE id = ?", 1).Error
	require.NoError(t, err)

	err = db.Exec(
		"INSERT INTO events (id, name) VALUES (?, ?)",
		1, "Integration Test Event",
	).Error
	require.NoError(t, err)

	err = db.Exec(
		`INSERT INTO seats
            (id, event_id, seat_number, status, version)
         VALUES (?, ?, ?, ?, ?)`,
		1, 1, "A1", "AVAILABLE", 0,
	).Error
	require.NoError(t, err)
}
