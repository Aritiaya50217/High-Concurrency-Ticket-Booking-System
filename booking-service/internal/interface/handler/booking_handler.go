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

	booking, err := h.usecase.Create(ctx, userID, req.EventID, req.SeatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.CreateBookingResponse{
		ID:     booking.ID,
		Status: string(booking.Status),
	}

	c.JSON(http.StatusCreated, res)
}

// func (h *BookingHandler) UpdateBooking(c *gin.Context) {
// 	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)

// 	var req dto.UpdateBookingRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
// 		return
// 	}

// 	if !utils.IsValidateStatus(req.Status) {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking status"})
// 		return
// 	}

// 	booking, err := h.usecase.Update(uint(id), req.Status)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	res := dto.UpdateBookingResponse{
// 		ID:     booking.ID,
// 		UserID: booking.UserID,
// 		Status: string(booking.Status),
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// func (h *BookingHandler) DeleteBooking(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 		return
// 	}

// 	if err := h.usecase.Delete(uint(id)); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "success"})
// }

// func (h *BookingHandler) FindBookingByID(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 		return
// 	}

// 	booking, err := h.usecase.FindByID(uint(id))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, booking)
// }

// func (h *BookingHandler) FindAll(c *gin.Context) {
// 	var req dto.SearchBookingRequest
// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
// 		return
// 	}

// 	// default values
// 	if req.Offset <= 0 {
// 		req.Offset = 1
// 	}

// 	if req.Limit <= 0 {
// 		req.Limit = 10
// 	}

// 	offset := (req.Offset - 1) * req.Limit

// 	bookings, total, err := h.usecase.Search(req.Status, offset, req.Limit)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result := []dto.SearchBookingResponse{}
// 	for _, booking := range bookings {
// 		result = append(result, dto.SearchBookingResponse{
// 			ID:     booking.ID,
// 			UserID: booking.UserID,
// 			// EventID: booking.EventID,
// 			SeatID: booking.SeatID,
// 			Status: string(booking.Status),
// 		})
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": result,
// 		"pagination": gin.H{
// 			"offset": req.Offset,
// 			"limit":  req.Limit,
// 			"total":  total,
// 		},
// 	})

// }
