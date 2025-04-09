package repository_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/okyws/dashboard-backend/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

type MockUserRepositoryAdapter struct {
	MockUserRepository *mock.Mock
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		Mock: mock.Mock{},
	}
}

func (m *MockUserRepositoryAdapter) GetUserByEmail(email string) (*domain.User, error) {
	args := m.MockUserRepository.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepositoryAdapter) GetUserByUsername(username string) (*domain.User, error) {
	args := m.MockUserRepository.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepositoryAdapter) Create(user *domain.User) (*domain.User, error) {
	args := m.MockUserRepository.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepositoryAdapter) Update(user *domain.User) (*domain.User, error) {
	args := m.MockUserRepository.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepositoryAdapter) Delete(id string) error {
	args := m.MockUserRepository.Called(id)

	return args.Error(0)
}

func TestGetUserByUsername(t *testing.T) {
	t.Run("User Found", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		id := uuid.New()
		mockUserRepository.On("GetUserByUsername", "username").Return(&domain.User{ID: id}, nil)

		user, err := mockUserRepositoryAdapter.GetUserByUsername("username")

		assert.NoError(t, err)
		assert.Equal(t, id, user.ID)
		mockUserRepository.AssertCalled(t, "GetUserByUsername", "username")
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		mockUserRepository.On("GetUserByUsername", "username").Return(nil, gorm.ErrRecordNotFound)

		user, err := mockUserRepositoryAdapter.GetUserByUsername("username")

		assert.Nil(t, user)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		mockUserRepository.AssertCalled(t, "GetUserByUsername", "username")
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		mockUserRepository.On("GetUserByUsername", "username").Return(nil, errors.New("database error"))

		user, err := mockUserRepositoryAdapter.GetUserByUsername("username")

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		mockUserRepository.AssertCalled(t, "GetUserByUsername", "username")
		mockUserRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("User Created", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		id := uuid.New()
		mockUserRepository.On("Create", mock.Anything).Return(&domain.User{ID: id}, nil)

		user, err := mockUserRepositoryAdapter.Create(&domain.User{})

		assert.NoError(t, err)
		assert.Equal(t, id, user.ID)
		mockUserRepository.AssertCalled(t, "Create", mock.Anything)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		mockUserRepository.On("Create", mock.Anything).Return(nil, errors.New("database error"))

		user, err := mockUserRepositoryAdapter.Create(&domain.User{})

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		mockUserRepository.AssertCalled(t, "Create", mock.Anything)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("Validation Error", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		mockUserRepository.On("Create", mock.Anything).Return(nil, errors.New("validation error"))

		user, err := mockUserRepositoryAdapter.Create(&domain.User{})

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, "validation error", err.Error())
		mockUserRepository.AssertCalled(t, "Create", mock.Anything)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("User Already Exists", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		mockUserRepository.On("Create", mock.Anything).Return(nil, errors.New("user already exists"))

		user, err := mockUserRepositoryAdapter.Create(&domain.User{})

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, "user already exists", err.Error())
		mockUserRepository.AssertCalled(t, "Create", mock.Anything)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		mockUserRepository.On("Create", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

		user, err := mockUserRepositoryAdapter.Create(&domain.User{})

		assert.Nil(t, user)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		mockUserRepository.AssertCalled(t, "Create", mock.Anything)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestMockUserRepositoryAdapterUpdate(t *testing.T) {
	t.Run("Update with valid user data", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		user := &domain.User{Username: "username", Email: "email"}
		mockUserRepository.On("Update", user).Return(user, nil)

		updatedUser, err := mockUserRepositoryAdapter.Update(user)

		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, user.Username, updatedUser.Username)
		assert.Equal(t, user.Email, updatedUser.Email)
		assert.Equal(t, user, updatedUser)
		mockUserRepository.AssertCalled(t, "Update", user)
	})

	t.Run("Update with invalid user data (nil)", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		var user *domain.User

		mockUserRepository.On("Update", user).Return(nil, errors.New("invalid user data"))

		updatedUser, err := mockUserRepositoryAdapter.Update(user)

		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		mockUserRepository.AssertCalled(t, "Update", user)
	})

	t.Run("Update with error from MockUserRepository", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		user := &domain.User{Username: "username", Email: "email"}
		mockUserRepository.On("Update", user).Return(nil, errors.New("database error"))

		updatedUser, err := mockUserRepositoryAdapter.Update(user)

		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		mockUserRepository.AssertCalled(t, "Update", user)
	})
}

func TestMockUserRepositoryAdapterDelete(t *testing.T) {
	t.Run("Delete with valid id and no error", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		id := "valid-id"
		mockUserRepository.On("Delete", id).Return(nil)

		err := mockUserRepositoryAdapter.Delete(id)

		assert.NoError(t, err)
		mockUserRepository.AssertCalled(t, "Delete", id)
	})

	t.Run("Delete with valid id and error", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		id := "valid-id"
		mockUserRepository.On("Delete", id).Return(errors.New("database error"))

		err := mockUserRepositoryAdapter.Delete(id)

		assert.Error(t, err)
		mockUserRepository.AssertCalled(t, "Delete", id)
	})

	t.Run("Delete with empty id and error", func(t *testing.T) {
		mockUserRepository := NewMockUserRepository()
		mockUserRepositoryAdapter := &MockUserRepositoryAdapter{MockUserRepository: &mockUserRepository.Mock}

		id := ""
		mockUserRepository.On("Delete", id).Return(errors.New("invalid id"))

		err := mockUserRepositoryAdapter.Delete(id)

		assert.Error(t, err)
		mockUserRepository.AssertCalled(t, "Delete", id)
	})
}
