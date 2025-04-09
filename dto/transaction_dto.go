package dto

import "github.com/google/uuid"

// TransactionDTO represents the transaction data transfer object for the API
type TransactionDTO struct {
	ID                uuid.UUID `json:"id"`
	FromAccountNumber string    `json:"from_account_number"`
	ToAccountNumber   string    `json:"to_account_number"`
	Amount            float64   `json:"amount"`
	TransactionType   string    `json:"transaction_type"`
	Status            string    `json:"status"`
}

// TransactionCreateDTO represents the transaction data transfer object for the API
type TransactionCreateDTO struct {
	FromAccountNumber string  `json:"from_account_number,omitempty" binding:"required_if=TransactionType transfer,required_if=TransactionType withdraw"`
	ToAccountNumber   string  `json:"to_account_number,omitempty" binding:"required_if=TransactionType transfer,required_if=TransactionType deposit"`
	Amount            float64 `json:"amount" binding:"required,min=10000"`
	TransactionType   string  `json:"transaction_type" binding:"required,oneof=deposit withdraw transfer"`
}
