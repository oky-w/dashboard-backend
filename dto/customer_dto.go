// Package dto contains the data transfer objects for the application
package dto

import (
	"time"

	"github.com/google/uuid"
)

// CustomerDTO represents the customer data transfer object for the API
type CustomerDTO struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address"`
	UserDTO     UserDTO   `json:"user,omitempty"`
}

// CustomerCreateDTO represents the customer data transfer object for the API
type CustomerCreateDTO struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	FullName    string    `json:"full_name" binding:"required,min=3,max=100"`
	PhoneNumber string    `json:"phone_number" binding:"required,e164"`
	DateOfBirth string    `json:"date_of_birth" binding:"required" time_format:"2006-01-02"`
	Address     string    `json:"address" binding:"omitempty,max=255"`
}

// CustomerUpdateDTO represents the customer data transfer object for the API
type CustomerUpdateDTO struct {
	FullName    string `json:"full_name" binding:"required,min=3,max=100"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	DateOfBirth string `json:"date_of_birth" binding:"required" time_format:"2006-01-02"`
	Address     string `json:"address" binding:"omitempty,max=255"`
}
