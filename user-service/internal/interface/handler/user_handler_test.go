package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/application/usecase"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/entity"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/security"
	"github.com/gin-gonic/gin"
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

func TestHandlerLoginSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &MockUserRepo{
		user: &entity.Users{
			ID:       1,
			Email:    "test@gmail.com",
			Password: hashPassword("123456"),
		},
	}
	jwt := security.NewJWTService("secret", time.Hour)
	u := usecase.NewUserUsecase(repo, jwt)
	handler := NewUserHandler(u, jwt)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"email":"test@gmail.com","password":"123456"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")

}
