// Package dto contains the data transfer objects for the application
package dto

import "github.com/google/uuid"

// BankAccountDTO represents the bank information data transfer object for the API
type BankAccountDTO struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	AccountType   string    `json:"account_type"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	AccountStatus bool      `json:"account_status"`
	UserDTO       UserDTO   `json:"user,omitempty"`
}

// BankAccountCreateDTO represents the bank information data transfer object for the API
type BankAccountCreateDTO struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	AccountType string    `json:"account_type" binding:"required,oneof=saku celengan deposito rekening-utama"`
	Balance     float64   `json:"balance" binding:"required,gte=0"`
}

// BankAccountUpdateDTO represents the bank information data transfer object for the API
type BankAccountUpdateDTO struct {
	Balance       float64 `json:"balance,omitempty" binding:"required,gte=0"`
	AccountStatus bool    `json:"account_status,omitempty" binding:"oneof=true false"`
}

// BankAccountUpdateStatusDTO represents the bank information data transfer object for the API
type BankAccountUpdateStatusDTO struct {
	AccountStatus bool `json:"account_status" binding:"required,oneof=true false"`
}
