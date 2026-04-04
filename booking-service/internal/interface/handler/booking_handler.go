package handler

import (
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/dto"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/pkg/utils"
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
		Status: booking.Status,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *BookingHandler) UpdateBooking(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	var req dto.UpdateBookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if !utils.IsValidateStatus(req.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking status"})
		return
	}

	booking, err := h.usecase.Update(uint(id), req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.UpdateBookingResponse{
		ID:     booking.ID,
		UserID: booking.UserID,
		Status: booking.Status,
	}

	c.JSON(http.StatusOK, res)
}

func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *BookingHandler) FindBookingByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	booking, err := h.usecase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, booking)
}
