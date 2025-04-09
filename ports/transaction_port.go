package ports

import "github.com/okyws/dashboard-backend/domain"

// TransactionRepository is the interface for the transaction repository
type TransactionRepository interface {
	GetAll(limit, offset int) ([]domain.Transaction, error)
	GetByID(id string) (*domain.Transaction, error)
	GetByAccountNumber(accountID string) ([]domain.Transaction, error)
	Create(transaction *domain.Transaction) (*domain.Transaction, error)
}

// TransactionService is the interface for the transaction service
type TransactionService interface {
	GetAllTransactions(limit, offset int) ([]domain.Transaction, error)
	GetTransactionByID(id string) (*domain.Transaction, error)
	GetTransactionByAccountNumber(accountID string) ([]domain.Transaction, error)
	ProcessTransaction(fromAccountID, toAccountID, transactionType string, amount float64) error
}
