package repository

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/database/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.Users) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entity.Users, error) {
	var userModel model.UsersModel
	if err := r.db.Where("email = ?", email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	user := &entity.Users{
		ID:       userModel.ID,
		Email:    userModel.Email,
		Password: userModel.Password,
	}

	return user, nil
}

func (r *UserRepository) FindByID(id uint) (*entity.Users, error) {
	var userModel model.UsersModel
	if err := r.db.First(&userModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	user := &entity.Users{
		ID:    userModel.ID,
		Email: userModel.Email,
	}
	return user, nil
}
