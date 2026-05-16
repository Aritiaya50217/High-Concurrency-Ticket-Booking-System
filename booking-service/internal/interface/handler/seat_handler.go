package handler

import (
	"net/http"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/dto"
	"github.com/gin-gonic/gin"
)

type SeatHandler struct {
	usecase *usecase.SeatUsecase
}

func NewSeatHandler(usecase *usecase.SeatUsecase) *SeatHandler {
	return &SeatHandler{usecase: usecase}
}

func (h *SeatHandler) Create(c *gin.Context) {
	var req dto.SeatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	seat := entity.Seat{
		// EventID:    req.EventID,
		SeatNumber: req.SeatNumber,
		Status:     string(valueobject.SeatAvailable),
		Version:    1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := h.usecase.Create(c, &seat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, seat)

}
