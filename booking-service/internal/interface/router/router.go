package router

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.BookingHandler, jwtService *security.JWTService) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	// middleware
	api.Use(middleware.AuthMiddleware(jwtService))
	api.POST("/booking", h.CreateBooking)
	api.POST("/booking/:id", h.UpdateBooking)
	api.DELETE("/booking/:id", h.DeleteBooking)
	api.GET("/booking/:id", h.FindBookingByID)
	api.GET("/booking/search", h.FindAll)
	return r
}
