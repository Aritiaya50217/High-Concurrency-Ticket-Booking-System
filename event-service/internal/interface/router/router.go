package router

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/metrics"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/middleware"
	authMiddleware "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetRouter(eventHandeler *handler.EventHandler, jwtService *security.JWTService) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.PrometheusMiddleware())

	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(metrics.Registry, promhttp.HandlerOpts{})))

	api := r.Group("/api")

	seat := api.Group("/seat", authMiddleware.AuthMiddleware(jwtService))
	seat.POST("/create", eventHandeler.CreateSeats)
	seat.POST("/reserve", eventHandeler.ReserveSeat)
	seat.POST("/release", eventHandeler.ReleaseSeat)

	event := api.Group("/event")
	event.POST("/create", eventHandeler.CreateEvent)

	return r
}
