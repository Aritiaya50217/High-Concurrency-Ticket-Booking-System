package router

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(bookingHandler *handler.BookingHandler, jwtService *security.JWTService) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	// middleware
	booking := api.Group("/booking")
	booking.Use(middleware.AuthMiddleware(jwtService))
	booking.POST("/create", bookingHandler.CreateBooking)
	return r
}
