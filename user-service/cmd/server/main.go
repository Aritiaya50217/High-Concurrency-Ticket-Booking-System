package main

import (
	"fmt"
	"log"
	"os"

	repositoryUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	handlerUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/router"
	usecaseUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/usecase"
	"github.com/joho/godotenv"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/database"
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

	jwtService := security.NewJWTService(cfg.JWT.Secret, cfg.JWTExpireDuration())

	// connect to database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repositoryUser := repositoryUser.NewUserRepository(db)
	usecaseUser := usecaseUser.NewUserUsecase(repositoryUser, jwtService)
	handlerUser := handlerUser.NewUserHandler(usecaseUser, jwtService)

	router := router.SetupRouter(handlerUser, jwtService)

	if err := router.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		log.Fatal(err)
	}
}
