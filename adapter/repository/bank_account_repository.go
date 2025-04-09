// Package repository contains the adapters for the bank information repository
package repository

import (
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
	"gorm.io/gorm"
)

// BankAccountRepositoryAdapter is the adapter for the bank information repository
type BankAccountRepositoryAdapter struct {
	db *gorm.DB
}

// NewBankAccountRepositoryAdapter creates a new bank information repository adapter via dependency injection
func NewBankAccountRepositoryAdapter(db *gorm.DB) ports.BankAccountRepository {
	return &BankAccountRepositoryAdapter{db: db}
}

// Create inserts a new bank information into the database
func (r *BankAccountRepositoryAdapter) Create(entity *domain.BankAccount) (*domain.BankAccount, error) {
	var createdData domain.BankAccount

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(entity).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Preload("User").First(&createdData, "id = ?", entity.ID).Error
	})

	if err != nil {
		return nil, err
	}

	return &createdData, nil
}

// Delete removes a bank information by ID
func (r *BankAccountRepositoryAdapter) Delete(id string) error {
	result := r.db.Delete(&domain.BankAccount{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAll fetches all bank information data with pagination
func (r *BankAccountRepositoryAdapter) GetAll(limit int, offset int) ([]domain.BankAccount, error) {
	var result []domain.BankAccount
	err := r.db.Preload("User").Limit(limit).Offset(offset).Find(&result).Error

	return result, err
}

// GetByID fetches a bank information by ID and returns nil if not found
func (r *BankAccountRepositoryAdapter) GetByID(id string) (*domain.BankAccount, error) {
	var result domain.BankAccount

	err := r.db.Preload("User").First(&result, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &result, nil
}

// Update updates a specific bank information
func (r *BankAccountRepositoryAdapter) Update(entity *domain.BankAccount) (*domain.BankAccount, error) {
	var updatedData domain.BankAccount

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("User").First(&updatedData, "id = ?", entity.ID).Error; err != nil {
			return err
		}

		if err := tx.Model(&updatedData).Updates(entity).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &updatedData, nil
}

// GetByUserID returns the bank information for a user
func (r *BankAccountRepositoryAdapter) GetByUserID(userID string) ([]domain.BankAccount, error) {
	var BankAccount []domain.BankAccount

	err := r.db.Preload("User").Find(&BankAccount, "user_id = ?", userID).Error

	return BankAccount, err
}

// CountBankAccount returns the count of main bank accounts
func (r *BankAccountRepositoryAdapter) CountBankAccount(userID string, accountType string) (int64, error) {
	var count int64

	err := r.db.Model(domain.BankAccount{}).Where("account_type = ? AND user_id = ?", accountType, userID).Count(&count).Error

	return count, err
}

// GetByAccountNumber fetches a bank account by its account number
func (r *BankAccountRepositoryAdapter) GetByAccountNumber(accountNumber string) (*domain.BankAccount, error) {
	var account domain.BankAccount
	if err := r.db.Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
