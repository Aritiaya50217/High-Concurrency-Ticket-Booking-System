package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/database"
	repositoryBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/repository"
	handlerBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/router"
	usecaseBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	// โหลด .env
	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		envPath = ".env"
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.yaml"
	}

	godotenv.Load(envPath)

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("App : ", cfg.App.Port)

	// connect to database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repoBooking := repositoryBooking.NewBookingRepository(db)
	usecaseBooking := usecaseBooking.NewBookingUsecase(repoBooking)
	handlerBooking := handlerBooking.NewBookingHandler(usecaseBooking)

	router := router.SetupRouter(handlerBooking)

	router.Run(fmt.Sprintf(":%d", cfg.App.Port))
	if err := router.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		log.Fatal(err)
	}
}
