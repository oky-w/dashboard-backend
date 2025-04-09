// Package services contains the business logic for the bank information service
package services

import (
	"errors"

	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
	"gorm.io/gorm"
)

// AccountValidator is a struct responsible for validating bank accounts.
type AccountValidator struct {
	UserRepository     ports.UserRepository
	BankInfoRepository ports.BankAccountRepository
}

// NewAccountValidator creates a new AccountValidator instance.
func NewAccountValidator(userRepo ports.UserRepository, bankInfoRepo ports.BankAccountRepository) *AccountValidator {
	return &AccountValidator{UserRepository: userRepo, BankInfoRepository: bankInfoRepo}
}

// validateUser checks if the user ID exists
func (v *AccountValidator) validateUser(userID string) error {
	_, err := v.UserRepository.GetByID(userID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return errors.New("user ID does not exist")
	}

	return nil
}

// validateAccountType checks if the account type is valid
func (v *AccountValidator) validateAccountType(bankInfo *domain.BankAccount) error {
	switch bankInfo.AccountType {
	case "rekening-utama":
		return v.validateMainAccount(bankInfo.UserID.String())
	case "saku", "deposito":
		return v.validateSecondaryAccount(bankInfo.UserID.String(), bankInfo.AccountType)
	default:
		return errors.New("invalid account type")
	}
}

// validateMainAccount checks if the user already has a main bank account
func (v *AccountValidator) validateMainAccount(userID string) error {
	count, err := v.BankInfoRepository.CountBankAccount(userID, "rekening-utama")
	if err != nil {
		return err
	}

	if count >= 1 {
		return errors.New("user already has a main bank account")
	}

	return nil
}

// validateSecondaryAccount checks if the user has a main bank account
func (v *AccountValidator) validateSecondaryAccount(userID string, accountType string) error {
	count, err := v.BankInfoRepository.CountBankAccount(userID, "rekening-utama")
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("user must have a main bank account to create a " + accountType + " account")
	}

	if accountType == "saku" && count >= 8 {
		return errors.New("saku account limit reached")
	}

	if accountType == "deposito" && count <= 3 {
		return errors.New("deposito account limit reached")
	}

	return nil
}
