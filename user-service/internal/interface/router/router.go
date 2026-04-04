package router

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/handler"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.UserHandler, jwtService *security.JWTService) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.POST("/login", h.Login)

	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware(jwtService))
	// auth.GET("/profile",)
	return r
}
