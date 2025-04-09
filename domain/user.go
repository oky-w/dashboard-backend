// Package domain contains the user model and its methods.
package domain

import (
	"github.com/google/uuid"
	"github.com/okyws/dashboard-backend/dto"
	"github.com/okyws/dashboard-backend/utils"
	"gorm.io/gorm"
)

// User struct represents the user model
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Email    string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Username string    `gorm:"type:varchar(50);unique;not null" json:"username"`
	Password string    `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Role     string    `gorm:"type:varchar(20);not null" json:"role"`
}

// BeforeCreate is a GORM hook to generate a UUID for the user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	if !utils.IsValidBcryptHash(u.Password) {
		hash, err := utils.GeneratePasswordHash(u.Password)
		if err != nil {
			tx.Rollback()
			return err
		}

		u.Password = hash
	}

	return nil
}

// MapUserToDTO maps a user to a UserDTO
func MapUserToDTO(user *User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}
}
