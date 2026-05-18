package router

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/interface/handler"
	authMiddleware "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/interface/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouter(eventHandeler *handler.EventHandler, jwtService *security.JWTService) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	seat := api.Group("/seat", authMiddleware.AuthMiddleware(jwtService))
	seat.POST("/create", eventHandeler.CreateSeats)
	seat.POST("/reserve", eventHandeler.ReserveSeat)
	seat.POST("/release", eventHandeler.ReleaseSeat)

	return r
}
