package mocks

import (
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *entity.Users) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(email string) (*entity.Users, error) {
	args := m.Called(email)

	var user *entity.Users

	if args.Get(0) != nil {
		user = args.Get(0).(*entity.Users)
	}
	return user, args.Error(1)
}

func (m *MockUserRepository) FindByID(id uint) (*entity.Users, error) {
	args := m.Called(id)

	var user *entity.Users

	if args.Get(0) != nil {
		user = args.Get(0).(*entity.Users)
	}

	return user, args.Error(1)
}

func (m *MockUserRepository) Profile(id uint) (*entity.Users, error) {
	args := m.Called(id)

	var user *entity.Users

	if args.Get(0) != nil {
		user = args.Get(0).(*entity.Users)
	}

	return user, args.Error(1)
}
