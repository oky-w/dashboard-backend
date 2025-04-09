package repository

import (
	"github.com/okyws/dashboard-backend/domain"
	"gorm.io/gorm"
)

// TransactionRepositoryAdapter is the adapter for the transaction repository
type TransactionRepositoryAdapter struct {
	db *gorm.DB
}

// NewTransactionRepositoryAdapter creates a new instance of TransactionRepositoryAdapter
func NewTransactionRepositoryAdapter(db *gorm.DB) *TransactionRepositoryAdapter {
	return &TransactionRepositoryAdapter{db: db}
}

// GetAll fetches all transactions with pagination
func (r *TransactionRepositoryAdapter) GetAll(limit, offset int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Limit(limit).Offset(offset).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetByID fetches a transaction by ID
func (r *TransactionRepositoryAdapter) GetByID(id string) (*domain.Transaction, error) {
	var transaction domain.Transaction
	if err := r.db.First(&transaction, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

// GetByAccountNumber fetches transactions by account number
func (r *TransactionRepositoryAdapter) GetByAccountNumber(accountID string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Where("from_account_number = ? OR to_account_number = ?", accountID, accountID).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

// Create adds a new transaction to the database
func (r *TransactionRepositoryAdapter) Create(transaction *domain.Transaction) (*domain.Transaction, error) {
	if err := r.db.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}
