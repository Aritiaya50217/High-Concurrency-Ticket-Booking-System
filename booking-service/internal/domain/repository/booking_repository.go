package repository

import "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/entity"

type BookingRepository interface {
	Create(booking *entity.Booking) error
	FindBySeatID(seatID uint) (*entity.Booking, error)
	FindBookingByID(id uint) (*entity.Booking, error)
	UpdateStatus(id uint, status string) error
}
