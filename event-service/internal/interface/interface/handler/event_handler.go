package handler

import (
	"net/http"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/interface/dto"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	usecase *usecase.EventUsecase
}

func NewEventHandler(usecase *usecase.EventUsecase) *EventHandler {
	return &EventHandler{usecase: usecase}
}

func (h *EventHandler) ReserveSeat(c *gin.Context) {
	var req dto.ReserveSeatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userID := c.GetUint("user_id")
	if err := h.usecase.ReserveSeat(c.Request.Context(), req.EventID, req.SeatID, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "seat resered"})

}
