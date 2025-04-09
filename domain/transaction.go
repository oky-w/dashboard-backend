// Package domain contains the transaction model
package domain

import (
	"github.com/google/uuid"
	"github.com/okyws/dashboard-backend/dto"
	"gorm.io/gorm"
)

// Transaction struct represents the transaction model
type Transaction struct {
	gorm.Model
	ID                uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	FromAccountNumber string    `gorm:"type:varchar(20);not null" json:"from_account_number"`
	ToAccountNumber   string    `gorm:"type:varchar(20);not null" json:"to_account_number"`
	Amount            float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	TransactionType   string    `gorm:"type:varchar(50);not null" json:"transaction_type"`
	Status            string    `gorm:"type:varchar(50);not null" json:"status"`
}

// BeforeCreate is a GORM hook to generate a UUID for the transaction
func (t *Transaction) BeforeCreate(_ *gorm.DB) error {
	t.ID = uuid.New()
	t.Status = "success"

	return nil
}

// MapTransactionToDTO maps a transaction to a transaction data transfer object
func MapTransactionToDTO(transaction *Transaction) *dto.TransactionDTO {
	return &dto.TransactionDTO{
		ID:                transaction.ID,
		FromAccountNumber: transaction.FromAccountNumber,
		ToAccountNumber:   transaction.ToAccountNumber,
		Amount:            transaction.Amount,
		TransactionType:   transaction.TransactionType,
		Status:            transaction.Status,
	}
}
