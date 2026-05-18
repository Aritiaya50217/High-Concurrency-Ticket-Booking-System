package dto

type CreateEventRequest struct {
	Name string `json:"name" binding:"required"`
}
