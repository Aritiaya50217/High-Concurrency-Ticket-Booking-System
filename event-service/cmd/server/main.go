package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/database"
	kafkaInfra "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/kafka"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/interface/router"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/worker"
	"github.com/joho/godotenv"

	infraRepo "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/repository"
)

func main() {

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

	eventRepo := infraRepo.NewEventRepository(db)

	inboxRepo := infraRepo.NewInboxRepository(db)

	seatUsecase := usecase.NewSeatUsecase(eventRepo)

	producer := kafkaInfra.NewProducer([]string{"localhost:9092"}, "seat-events")

	eventProducerRepo := infraRepo.NewEventProducerRepository(producer)

	eventUsecase := usecase.NewEventUsecase(eventRepo, inboxRepo, eventProducerRepo)

	eventHandler := handler.NewEventHandler(eventUsecase, seatUsecase)

	// ----------------
	// Kafka Consumer
	// ----------------
	consumer := kafkaInfra.NewConsumer(
		cfg.Kafka.Brokers,
		"booking-events",
		"event-service",
	)

	bookingCosumer := worker.NewBookingCreatedConsumer(consumer, eventUsecase)

	go bookingCosumer.Start()

	log.Println("Kafka consumer started")

	r := router.SetRouter(eventHandler, jwtService)

	// ----------------------------
	// Run server
	// ----------------------------
	log.Println("Server running...")
	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		log.Fatal(err)
	}
}
