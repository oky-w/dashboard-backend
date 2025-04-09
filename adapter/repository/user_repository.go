// Package repository contains the adapters for the user repository
package repository

import (
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
	"gorm.io/gorm"
)

// UserRepositoryAdapter is the adapter for the user repository
type UserRepositoryAdapter struct {
	db *gorm.DB
}

// NewUserRepositoryAdapter creates a new user repository adapter via dependency injection
func NewUserRepositoryAdapter(db *gorm.DB) ports.UserRepository {
	return &UserRepositoryAdapter{db: db}
}

// Create inserts a new user into the database using a transaction.
func (r *UserRepositoryAdapter) Create(user *domain.User) (*domain.User, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByID fetches a user by ID, returning nil if not found.
func (r *UserRepositoryAdapter) GetByID(id string) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

// GetUserByEmail fetches a user by email, returning nil if not found.
func (r *UserRepositoryAdapter) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

// GetUserByUsername fetches a user by username, returning nil if not found.
func (r *UserRepositoryAdapter) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

// Update updates an existing user in the database.
func (r *UserRepositoryAdapter) Update(user *domain.User) (*domain.User, error) {
	var updatedUser domain.User

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&updatedUser, "id = ?", user.ID).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Model(&updatedUser).Updates(user).Error
	})

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

// Delete removes a user by ID and ensures that a user was actually deleted.
func (r *UserRepositoryAdapter) Delete(id string) error {
	result := r.db.Delete(&domain.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAll fetches users with pagination.
func (r *UserRepositoryAdapter) GetAll(limit, offset int) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error

	return users, err
}
