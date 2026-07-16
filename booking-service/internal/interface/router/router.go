package router

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/metrics"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(bookingHandler *handler.BookingHandler, jwtService *security.JWTService) *gin.Engine {
	r := gin.Default()

	// Collect HTTP metrics
	r.Use(middleware.PrometheusMiddleware())

	// Prometheus Metrics
	r.GET("/metrics", gin.WrapH(
		promhttp.HandlerFor(metrics.Registery, promhttp.HandlerOpts{}),
	))

	api := r.Group("/api")

	// middleware
	booking := api.Group("/booking")
	booking.Use(middleware.AuthMiddleware(jwtService))
	booking.POST("/create", bookingHandler.CreateBooking)
	
	return r
}
