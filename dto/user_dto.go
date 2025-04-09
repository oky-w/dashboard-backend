// Package dto contains the data transfer objects for the application.
package dto

import "github.com/google/uuid"

// UserDTO represents the user data transfer object for the API
type UserDTO struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Email    string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Username string    `gorm:"type:varchar(50);unique;not null" json:"username"`
	Role     string    `gorm:"type:varchar(20);not null" json:"role"`
}

// UserCreateDTO represents the user data transfer object for the API
type UserCreateDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin user customer"`
}

// UserUpdateDTO represents the user data transfer object for the API
type UserUpdateDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password,omitempty" binding:"omitempty,min=6"`
	Role     string `json:"role,omitempty" binding:"omitempty,oneof=admin user customer"`
}

// UserLoginDTO represents the user data transfer object for the API
type UserLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserLoginResponseDTO represents the user login response data transfer object for the API
type UserLoginResponseDTO struct {
	Username  string `json:"username"`
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}
