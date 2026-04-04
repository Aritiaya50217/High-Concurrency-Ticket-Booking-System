package router

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.BookingHandler) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.POST("/booking", h.CreateBooking)
	return r
}
