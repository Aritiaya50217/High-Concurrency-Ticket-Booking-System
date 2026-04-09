package repository

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.Users) error
	FindByEmail(email string) (*entity.Users, error)
	FindByID(id uint) (*entity.Users, error)
}
