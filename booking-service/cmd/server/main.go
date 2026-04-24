package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/database"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/kafka"
	repositoryBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/repository"
	handlerBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/router"
	usecaseBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/usecase"
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
		cfg.Kafka.TopicBookingCreated,
		cfg.Kafka.GroupID,
	)

	// ----------------------------
	// Repository
	// ----------------------------
	repoBooking := repositoryBooking.NewBookingRepository(db)
	repoSeat := repositoryBooking.NewSeatRepository(db)

	// ----------------------------
	// Usecase
	// ----------------------------
	bookingUsecase := usecaseBooking.NewBookingUsecase(
		repoBooking,
		repoSeat,
		producer,
		cfg.Kafka.TopicBookingCreated,
	)

	// ----------------------------
	// Handler
	// ----------------------------
	bookingHandler := handlerBooking.NewBookingHandler(bookingUsecase)

	// ----------------------------
	// Router
	// ----------------------------
	r := router.SetupRouter(bookingHandler)

	// ----------------------------
	// Worker (Kafka Consumer)
	// ----------------------------
	bookingWorker := worker.NewBookingConsumer(consumer)
	go bookingWorker.Start(ctx)

	// ----------------------------
	// Run server
	// ----------------------------
	log.Println("🌐 Server running...")
	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		log.Fatal(err)
	}
}
