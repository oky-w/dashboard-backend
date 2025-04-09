// Package services contains the business logic for the bank information service
package services

import (
	"errors"

	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
)

// BankAccountService is the implementation of the bank information service
type BankAccountService struct {
	UserRepository     ports.UserRepository
	BankInfoRepository ports.BankAccountRepository
	AccountValidator   *AccountValidator
}

// NewBankAccountService creates a new bank information service
func NewBankAccountService(userRepo ports.UserRepository, bankInfoRepo ports.BankAccountRepository, validator *AccountValidator) *BankAccountService {
	return &BankAccountService{UserRepository: userRepo, BankInfoRepository: bankInfoRepo, AccountValidator: validator}
}

// CreateBankAccount creates a new bank information
func (s *BankAccountService) CreateBankAccount(bankInfo *domain.BankAccount) (*domain.BankAccount, error) {
	if err := s.AccountValidator.validateUser(bankInfo.UserID.String()); err != nil {
		return nil, err
	}

	if err := s.AccountValidator.validateAccountType(bankInfo); err != nil {
		return nil, err
	}

	return s.BankInfoRepository.Create(bankInfo)
}

// UpdateBankAccount updates a specific bank information
func (s *BankAccountService) UpdateBankAccount(bankInfo *domain.BankAccount) (*domain.BankAccount, error) {
	return s.BankInfoRepository.Update(bankInfo)
}

// DeleteBankAccount deletes a specific bank information
func (s *BankAccountService) DeleteBankAccount(id string) error {
	exist, err := s.GetBankAccountByID(id)
	if err != nil {
		return err
	}

	if exist == nil {
		return errors.New("bank information not found")
	}

	if exist.AccountType == "rekening-utama" {
		return errors.New("contact admin to delete main bank account")
	}

	return s.BankInfoRepository.Delete(id)
}

// GetBankAccountByID fetches a bank information by ID and returns nil if not found
func (s *BankAccountService) GetBankAccountByID(id string) (*domain.BankAccount, error) {
	return s.BankInfoRepository.GetByID(id)
}

// GetAllBankAccount fetches all bank information data with pagination
func (s *BankAccountService) GetAllBankAccount(limit int, offset int) ([]domain.BankAccount, error) {
	return s.BankInfoRepository.GetAll(limit, offset)
}

// GetByUserID returns the bank information for a user
func (s *BankAccountService) GetByUserID(userID string) ([]domain.BankAccount, error) {
	return s.BankInfoRepository.GetByUserID(userID)
}
