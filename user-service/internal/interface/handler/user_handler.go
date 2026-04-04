package handler

import (
	"net/http"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/dto"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase    *usecase.UserUsecase
	jwtService *security.JWTService
}

func NewUserHandler(usecase *usecase.UserUsecase, jwtService *security.JWTService) *UserHandler {
	return &UserHandler{usecase: usecase, jwtService: jwtService}
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
