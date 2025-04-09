// Package services contains the business logic for the user service
package services

import (
	"errors"

	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
	"github.com/okyws/dashboard-backend/utils"
	"gorm.io/gorm"
)

// UserService is the implementation of the user service
type UserService struct {
	UserRepository ports.UserRepository
}

// NewUserService creates a new user service via dependency injection
func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

// CreateUser inserts a new user into the database
func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	if user.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	existingUser, err := s.GetUserByUsername(user.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	createdUser, err := s.UserRepository.Create(user)
	if err != nil || createdUser == nil {
		return nil, err
	}

	createdUser.Password = ""

	return createdUser, nil
}

// GetUserByID fetches a user by ID
func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	user, err := s.UserRepository.GetByID(id)
	if err != nil || user == nil {
		return nil, err
	}

	user.Password = ""

	return user, err
}

// GetUserByUsername fetches a user by username
func (s *UserService) GetUserByUsername(username string) (*domain.User, error) {
	user, err := s.UserRepository.GetUserByUsername(username)
	if err != nil || user == nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *domain.User) (*domain.User, error) {
	if user.Password != "" {
		hash, err := utils.GeneratePasswordHash(user.Password)
		if err != nil {
			return nil, err
		}

		user.Password = hash
	}

	updatedUser, err := s.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}

	updatedUser.Password = ""

	return updatedUser, nil
}

// DeleteUser removes a user
func (s *UserService) DeleteUser(id string) error {
	return s.UserRepository.Delete(id)
}

// GetAllUsers fetches all users
func (s *UserService) GetAllUsers(limit, offset int) ([]domain.User, error) {
	return s.UserRepository.GetAll(limit, offset)
}
