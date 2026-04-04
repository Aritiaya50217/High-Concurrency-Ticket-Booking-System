package main

import (
	"fmt"
	"log"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/database"
	repositoryBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/repository"
	handlerBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/router"
	usecaseBooking "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("App : ", cfg.App.Port)

	// connect to database
	db := database.NewPostgresDB(cfg)

	repoBooking := repositoryBooking.NewBookingRepository(db)
	usecaseBooking := usecaseBooking.NewBookingUsecase(repoBooking)
	handlerBooking := handlerBooking.NewBookingHandler(usecaseBooking)

	router := router.SetupRouter(handlerBooking)

	router.Run(fmt.Sprintf(":%d", cfg.App.Port))
}
