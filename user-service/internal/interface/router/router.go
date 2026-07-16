package router

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/metrics"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(h *handler.UserHandler, jwtService *security.JWTService) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.PrometheusMiddleware())

	r.GET("/metrics", gin.WrapH(
		promhttp.HandlerFor(metrics.Registry, promhttp.HandlerOpts{}),
	))

	api := r.Group("/api")
	{
		api.POST("/login", h.Login)
		api.POST("/register", h.Register)

		// middleware
		api.Use(middleware.AuthMiddleware(jwtService))
		api.GET("/profile/", h.Profile)
	}
	return r
}
