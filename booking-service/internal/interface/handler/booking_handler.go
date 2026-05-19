package handler

import (
	"context"
	"net/http"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/dto"
	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	usecase *usecase.BookingUsecase
}

func NewBookingHandler(usecase *usecase.BookingUsecase) *BookingHandler {
	return &BookingHandler{usecase: usecase}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req dto.CreateBookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

	userID := user.(uint)

	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized,
			gin.H{"error": "missing token"})
		return
	}

	ctx := context.WithValue(c.Request.Context(), "token", token)

	_, err := h.usecase.Create(ctx, userID, req.EventID, req.SeatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}
