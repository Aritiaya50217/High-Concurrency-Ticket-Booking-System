package usecase

import (
	"errors"
	"fmt"

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

// func (u *UserUsecase) Register(email, password string) error {

// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	user := entity.Users{
// 		Email:    email,
// 		Password: string(hash),
// 	}

// 	return u.repo.Create(&user)
// }

func (u *UserUsecase) Login(email, password string) (*entity.Users, error) {

	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	fmt.Println("DB password:", user.Password)
	fmt.Println("Input password:", password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("bcrypt error:", err)
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (u *UserUsecase) ValidateToken(tokenStr string) (uint, error) {
	return u.jwtService.ValidateToken(tokenStr)
}
