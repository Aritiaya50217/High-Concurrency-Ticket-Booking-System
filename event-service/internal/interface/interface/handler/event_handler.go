package handler

import (
	"net/http"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/interface/interface/dto"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventUsecase *usecase.EventUsecase
	seatUsecase  *usecase.SeatUsecase
}

func NewEventHandler(eventUsecase *usecase.EventUsecase, seatUsecase *usecase.SeatUsecase) *EventHandler {
	return &EventHandler{eventUsecase: eventUsecase, seatUsecase: seatUsecase}
}

func (h *EventHandler) ReserveSeat(c *gin.Context) {
	var req dto.ReserveSeatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.seatUsecase.Reserve(c.Request.Context(), req.EventID, req.SeatID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "seat resered"})

}

func (h *EventHandler) ReleaseSeat(c *gin.Context) {
	var req dto.ReleaseSeatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.seatUsecase.Release(c.Request.Context(), req.EventID, req.SeatID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "seat resered"})

}
