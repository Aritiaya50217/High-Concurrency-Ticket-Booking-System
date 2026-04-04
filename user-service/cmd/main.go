package main

import (
	"fmt"
	"log"

	repositoryUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	handlerUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/router"
	usecaseUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/usecase"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/database"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("App : ", cfg.App.Port)

	jwtService := security.NewJWTService(cfg.JWT.Secret, cfg.JWTExpireDuration())

	// connect to database
	db := database.NewPostgresDB(cfg)

	repositoryUser := repositoryUser.NewUserRepository(db)
	usecaseUser := usecaseUser.NewUserUsecase(repositoryUser, jwtService)
	handlerUser := handlerUser.NewUserHandler(usecaseUser, jwtService)

	router := router.SetupRouter(handlerUser, jwtService)

	router.Run(fmt.Sprintf(":%d", cfg.App.Port))
}
