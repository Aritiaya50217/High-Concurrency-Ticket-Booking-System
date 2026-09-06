package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/database"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/external/eventservice"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/grpc"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/router"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/worker"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	// โหลด .env
	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		envPath = ".env"
	}
	_ = godotenv.Load(envPath)

	// โหลด config
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.yaml"
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("App Port:", cfg.App.Port)

	jwtService := security.NewJWTService(cfg.JWT.Secret, cfg.JWTExpireDuration())

	// ----------------------------
	// Database
	// ----------------------------
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// ----------------------------
	// Kafka
	// ----------------------------
	producer := kafka.NewProducer(cfg.Kafka.Brokers)

	consumer := kafka.NewConsumer(
		cfg.Kafka.Brokers,
		cfg.Kafka.TopicSeatReserved,
		cfg.Kafka.GroupID,
	)

	// ----------------------------
	// Repository
	// ----------------------------
	repoBooking := repository.NewBookingRepository(db)

	repoOutbox := repository.NewOutboxRepository(db)

	// ----------------------------
	// External
	// ----------------------------
	baseEventURL := os.Getenv("BASE_EVENT_URL")
	if baseEventURL == "" {
		baseEventURL = "http://localhost:8082"
	}
	eventClient := eventservice.NewSeatServiceClient(baseEventURL)

	// ----------------------------
	// gRPC
	// ----------------------------
	userClient, err := grpc.NewUserServiceClient(cfg.GRPC.UserServiceAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer userClient.Close()

	exists, err := userClient.GetUser(
		context.Background(),
		1,
	)

	if err != nil {
		log.Printf("gRPC GetUser error: %v", err)
	} else {
		log.Printf("gRPC GetUser result: exists=%v", exists)
	}

	// ----------------------------
	// Usecase
	// ----------------------------
	bookingUsecase := usecase.NewBookingUsecase(
		repoBooking,
		producer,
		cfg.Kafka.TopicBookingCreated,
		*eventClient,
		userClient,
	)

	// ----------------------------
	// Handler
	// ----------------------------
	bookingHandler := handler.NewBookingHandler(bookingUsecase)

	// ----------------------------
	// Router
	// ----------------------------
	r := router.SetupRouter(bookingHandler, jwtService)

	// ----------------------------
	// Worker (Kafka Producer)
	// ----------------------------
	// polling outbox -> publish kafka
	outboxWorker := worker.NewOutboxWorker(repoOutbox, producer, cfg.Kafka.TopicBookingCreated)

	go outboxWorker.Start(ctx)

	// ----------------------------
	// Worker (Kafka Consumer)
	// ----------------------------
	bookingWorker := worker.NewBookingConsumer(consumer, bookingUsecase)
	go bookingWorker.Start(ctx)

	// ----------------------------
	// Expiration Worker
	// ----------------------------
	expirationWorker := worker.NewBookingExpirationWorker(repoBooking, repoOutbox)

	go expirationWorker.Start()

	// ----------------------------
	// Run server
	// ----------------------------
	log.Println("Server running...")
	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		log.Fatal(err)
	}
}
