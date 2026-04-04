package usecase

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo       repository.UserRepository
	jwtService *security.JWTService
}

func NewUserUsecase(repo repository.UserRepository, jwtService *security.JWTService) *UserUsecase {
	return &UserUsecase{repo: repo, jwtService: jwtService}
}

func (u *UserUsecase) Login(email, password string) (*entity.Users, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func (u *UserUsecase) ValidateToken(tokenStr string) (uint, error) {
	return u.jwtService.ValidateToken(tokenStr)
}
