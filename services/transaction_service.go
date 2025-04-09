package services

import (
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
	"gorm.io/gorm"
)

// TransactionService is the implementation of the transaction service
type TransactionService struct {
	db                    *gorm.DB
	TransactionRepository ports.TransactionRepository
	BankInfoRepository    ports.BankAccountRepository
	TransactionValidator  *TransactionValidator
}

// NewTransactionService creates a new transaction service
func NewTransactionService(db *gorm.DB, transactionRepo ports.TransactionRepository, bankInfoRepo ports.BankAccountRepository, validator *TransactionValidator) *TransactionService {
	return &TransactionService{
		db:                    db,
		TransactionRepository: transactionRepo,
		BankInfoRepository:    bankInfoRepo,
		TransactionValidator:  validator,
	}
}

// ProcessTransaction processes a transaction based on its type
func (s *TransactionService) ProcessTransaction(fromAccountNumber, toAccountNumber, transactionType string, amount float64) error {
	err := s.db.Transaction(func(_ *gorm.DB) error {
		transaction := domain.Transaction{
			FromAccountNumber: fromAccountNumber,
			ToAccountNumber:   toAccountNumber,
			Amount:            amount,
			TransactionType:   transactionType,
		}

		if err := s.TransactionValidator.ProcessTransaction(&transaction); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// GetAllTransactions retrieves all transactions with pagination
func (s *TransactionService) GetAllTransactions(limit, offset int) ([]domain.Transaction, error) {
	return s.TransactionRepository.GetAll(limit, offset)
}

// GetTransactionByID retrieves a specific transaction by its ID
func (s *TransactionService) GetTransactionByID(id string) (*domain.Transaction, error) {
	return s.TransactionRepository.GetByID(id)
}

// GetTransactionByAccountID retrieves transactions by account ID
func (s *TransactionService) GetTransactionByAccountID(accountID string) ([]domain.Transaction, error) {
	return s.TransactionRepository.GetByAccountNumber(accountID)
}
