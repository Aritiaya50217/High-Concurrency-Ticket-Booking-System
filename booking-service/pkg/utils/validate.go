package utils

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
)

func IsValidateStatus(status string) bool {
	switch status {
	case string(valueobject.BookingConfirmed), string(valueobject.BookingPending), string(valueobject.BookingCancelled):
		return true
	default:
		return false
	}

}
