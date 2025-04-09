// Package domain contains the bank information model
package domain

import (
	"github.com/google/uuid"
	"github.com/okyws/dashboard-backend/dto"
	"github.com/okyws/dashboard-backend/utils"
	"gorm.io/gorm"
)

// BankAccount struct represents the bank information model
type BankAccount struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User          *User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"user"` // Relasi ke User
	AccountType   string    `gorm:"type:varchar(255);not null" json:"account_type"`
	AccountNumber string    `gorm:"type:varchar(255);unique;not null" json:"account_number"`
	Balance       float64   `gorm:"type:decimal(10,2);not null;default:0" json:"last_balance"`
	AccountStatus bool      `gorm:"type:bool;not null" json:"account_status"`
}

// BeforeCreate is a GORM hook to generate a UUID for the bank information
func (b *BankAccount) BeforeCreate(tx *gorm.DB) error {
	result, err := utils.GenerateAccountNumber(10)
	if err != nil || result == "" {
		tx.Rollback()
		return err
	}

	b.ID = uuid.New()
	b.AccountNumber = result
	b.AccountStatus = true

	return nil
}

// MapBankAccountToDTO maps a bank information to a BankAccountDTO
func MapBankAccountToDTO(bankAccount *BankAccount) *dto.BankAccountDTO {
	return &dto.BankAccountDTO{
		ID:            bankAccount.ID,
		UserID:        bankAccount.UserID,
		AccountType:   bankAccount.AccountType,
		AccountNumber: bankAccount.AccountNumber,
		Balance:       bankAccount.Balance,
		AccountStatus: bankAccount.AccountStatus,
		UserDTO:       *MapUserToDTO(bankAccount.User),
	}
}
