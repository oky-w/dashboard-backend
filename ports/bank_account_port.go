package ports

import (
	"github.com/okyws/dashboard-backend/domain"
)

// BankAccountRepository is the interface for the bank information repository
type BankAccountRepository interface {
	GenericRepository[domain.BankAccount]
	GetByUserID(userID string) ([]domain.BankAccount, error)
	GetByAccountNumber(accountNumber string) (*domain.BankAccount, error)
	CountBankAccount(userID string, accountType string) (int64, error)
}

// BankAccountService is the interface for the bank information service
type BankAccountService interface {
	GenericService[domain.BankAccount]
	GetByUserID(userID string) ([]domain.BankAccount, error)
}
