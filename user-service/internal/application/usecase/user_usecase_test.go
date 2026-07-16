package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepo struct {
	user *entity.Users
	err  error
}

func (m *MockUserRepo) FindByEmail(email string) (*entity.Users, error) {
	return m.user, m.err
}

func (m *MockUserRepo) Create(user *entity.Users) error {
	return nil
}

func (m *MockUserRepo) FindByID(id uint) (*entity.Users, error) {
	return nil, nil
}

func (m *MockUserRepo) Profile(id uint) (*entity.Users, error) {
	return nil, nil
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func TestLoginSuccess(t *testing.T) {
	repo := &MockUserRepo{
		user: &entity.Users{
			ID:       1,
			Email:    "test@gmail.com",
			Password: hashPassword("123456"),
		},
	}
	jwt := security.NewJWTService("secert", time.Hour)
	u := NewUserUsecase(repo, jwt)

	token, err := u.Login("test@gmail.com", "123456")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginUserNotFound(t *testing.T) {
	repo := &MockUserRepo{user: nil, err: nil}
	jwt := security.NewJWTService("secret", time.Hour)
	u := NewUserUsecase(repo, jwt)
	token, err := u.Login("no@gmail.com", "123456")

	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, "user not found", err.Error())

}

func TestLoginWrongPassword(t *testing.T) {
	repo := &MockUserRepo{
		user: &entity.Users{
			ID:       1,
			Email:    "test@gmail.com",
			Password: hashPassword("123456"),
		},
	}
	jwt := security.NewJWTService("secret", time.Hour)
	u := NewUserUsecase(repo, jwt)

	token, err := u.Login("test@gmail.com", "wrong")
	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, "invalid email or password", err.Error())
}

func TestLoginDBError(t *testing.T) {
	repo := &MockUserRepo{
		user: nil,
		err:  errors.New("db error"),
	}

	jwt := security.NewJWTService("secret", time.Hour)
	u := NewUserUsecase(repo, jwt)

	token, err := u.Login("test@gmail.com", "123456")

	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, "db error", err.Error())
}
