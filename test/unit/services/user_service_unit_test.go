package services_test

import (
	"errors"
	"testing"

	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *MockUserRepository) GetAll(limit, offset int) ([]domain.User, error) {
	args := m.Called(limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id string) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func TestCreateUserService(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := services.NewUserService(mockRepo)

	t.Run("Empty password", func(t *testing.T) {
		user := &domain.User{Username: "testuser", Password: ""}
		_, err := userService.CreateUser(user)
		assert.Error(t, err)
		assert.Equal(t, "password cannot be empty", err.Error())
	})

	t.Run("Existing username", func(t *testing.T) {
		existingUser := &domain.User{Username: "testuser", Password: "testpassword"}
		mockRepo.On("GetUserByUsername", existingUser.Username).Return(existingUser, nil)

		user := &domain.User{Username: "testuser", Password: "newpassword"}
		_, err := userService.CreateUser(user)

		assert.Error(t, err)
		assert.Equal(t, "username already exists", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Successful user creation", func(t *testing.T) {
		user := &domain.User{Username: "newuser", Password: "newpassword"}

		mockRepo.On("GetUserByUsername", "newuser").Return(nil, nil)
		mockRepo.On("Create", user).Return(user, nil)

		createdUser, err := userService.CreateUser(user)

		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, "newuser", createdUser.Username)
		assert.Equal(t, "", createdUser.Password)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Database error", func(t *testing.T) {
		user := &domain.User{Username: "newuser", Password: "newpassword"}

		mockRepo.On("GetUserByUsername", "newuser").Return(nil, nil)
		mockRepo.On("Create", user).Return(nil, errors.New("database error"))

		createdUser, err := userService.CreateUser(user)

		assert.Error(t, err)
		assert.Nil(t, createdUser)
		assert.Equal(t, "database error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
