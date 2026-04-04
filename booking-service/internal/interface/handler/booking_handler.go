package handler

import (
	"net/http"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/dto"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/usecase"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking, err := h.usecase.Create(req.UserID, req.EventID, req.SeatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.CreateBookingResponse{
		ID:     booking.ID,
		Status: booking.Status,
	}

	c.JSON(http.StatusCreated, res)
}
