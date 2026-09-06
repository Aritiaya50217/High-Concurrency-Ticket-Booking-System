package main

import (
	"fmt"
	"log"
	"net"
	"os"

	repositoryUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	handlerUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/router"
	usecaseUser "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/usecase"
	"github.com/joho/godotenv"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/database"

	grpcServer "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcHandler "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/grpc"

	userpb "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/contracts/user"
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

	// grpc
	server := grpcServer.NewServer()
	userServer := grpcHandler.NewUserServer(usecaseUser)
	userpb.RegisterUserServiceServer(server, userServer)
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server running on :50051")

	// gRPC server
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// HTTP server
	router := router.SetupRouter(handlerUser, jwtService)
	if err := router.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		log.Fatal(err)
	}
}
