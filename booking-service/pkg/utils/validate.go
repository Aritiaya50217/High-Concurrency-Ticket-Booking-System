package utils

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"
)

func IsValidateStatus(status string) bool {
	switch status {
	case entity.StatusConfirmed, entity.StatusPending, entity.StatusCanceled:
		return true
	default:
		return false
	}

}
