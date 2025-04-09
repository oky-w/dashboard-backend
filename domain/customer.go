// Package domain contains the customer model and its methods
package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/okyws/dashboard-backend/dto"
	"gorm.io/gorm"
)

// Customer struct represents the customer model
type Customer struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;unique;not null" json:"user_id"`
	User        *User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"user"` // Relasi ke User
	FullName    string    `gorm:"type:varchar(100);not null" json:"full_name"`
	PhoneNumber string    `gorm:"type:varchar(20);unique;not null" json:"phone_number"`
	DateOfBirth time.Time `gorm:"type:date;not null" json:"date_of_birth" time_format:"2006-01-02"`
	Address     string    `gorm:"type:text" json:"address"`
}

// BeforeCreate is a GORM hook to generate a UUID for the customer
func (c *Customer) BeforeCreate(_ *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}

// MapCustomerToDTO maps a customer to a CustomerDTO
func MapCustomerToDTO(customer *Customer) *dto.CustomerDTO {
	return &dto.CustomerDTO{
		ID:          customer.ID,
		UserID:      customer.UserID,
		FullName:    customer.FullName,
		PhoneNumber: customer.PhoneNumber,
		DateOfBirth: customer.DateOfBirth,
		Address:     customer.Address,
		UserDTO:     *MapUserToDTO(customer.User),
	}
}
