package usecase

import (
	"errors"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/repository"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/metrics"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/interface/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo       repository.UserRepository
	jwtService *security.JWTService
}

func NewUserUsecase(repo repository.UserRepository, jwtService *security.JWTService) *UserUsecase {
	return &UserUsecase{repo: repo, jwtService: jwtService}
}

func (u *UserUsecase) Register(req dto.RegisterRequest) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.Users{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(hash),
	}

	if err := u.repo.Create(&user); err != nil {
		return err
	}

	metrics.UserRegisteredTotal.Inc()

	return nil
}

func (u *UserUsecase) Login(email, password string) (string, error) {

	user, err := u.repo.FindByEmail(email)
	if err != nil {
		metrics.UserLoginFailedTotal.WithLabelValues("failed").Inc()
		return "", err
	}

	if user == nil {
		metrics.UserLoginFailedTotal.WithLabelValues("failed").Inc()
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		metrics.UserLoginFailedTotal.WithLabelValues("failed").Inc()
		return "", errors.New("invalid email or password")
	}

	token, err := u.jwtService.GenerateToken(user.ID)
	if err != nil {
		metrics.UserLoginFailedTotal.WithLabelValues("failed").Inc()
		return "", err
	}

	metrics.UserLoginSuccessTotal.WithLabelValues("success").Inc()

	return token, nil
}
func (u *UserUsecase) ValidateToken(tokenStr string) (uint, error) {
	return u.jwtService.ValidateToken(tokenStr)
}

func (u *UserUsecase) Profile(id uint) (*entity.Users, error) {
	user, err := u.repo.Profile(id)
	if err != nil {
		metrics.UserGetProfileFailedTotal.Inc()
		return nil, err
	}

	metrics.UserGetProfileSuccessTotal.Inc()

	return user, nil
}
