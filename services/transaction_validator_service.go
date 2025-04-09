package services

import (
	"errors"

	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
)

// TransactionValidator is a struct responsible for validating transactions.
type TransactionValidator struct {
	TransactionRepository ports.TransactionRepository
	BankInfoRepository    ports.BankAccountRepository
}

// NewTransactionValidator creates a new TransactionValidator instance.
func NewTransactionValidator(transactionRepo ports.TransactionRepository, bankInfoRepo ports.BankAccountRepository) *TransactionValidator {
	return &TransactionValidator{TransactionRepository: transactionRepo, BankInfoRepository: bankInfoRepo}
}

// ProcessTransaction processes a transaction based on its type.
func (s *TransactionValidator) ProcessTransaction(transaction *domain.Transaction) error {
	switch transaction.TransactionType {
	case "transfer":
		return s.processTransfer(transaction.FromAccountNumber, transaction.ToAccountNumber, transaction.Amount)
	case "deposit":
		return s.processDeposit(transaction.ToAccountNumber, transaction.Amount)
	case "withdraw":
		return s.processWithdraw(transaction.FromAccountNumber, transaction.Amount)
	default:
		return errors.New("invalid transaction type")
	}
}

// Helper function to process a transfer transaction
func (s *TransactionValidator) processTransfer(fromAccountNumber, toAccountNumber string, amount float64) error {
	fromAccount, err := s.BankInfoRepository.GetByAccountNumber(fromAccountNumber)
	if err != nil {
		return err
	}

	toAccount, err := s.BankInfoRepository.GetByAccountNumber(toAccountNumber)
	if err != nil {
		return err
	}

	if fromAccount.AccountNumber == toAccount.AccountNumber {
		return errors.New("cannot transfer to the same account")
	}

	if fromAccount.Balance <= amount {
		return errors.New("insufficient balance")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	if _, err := s.BankInfoRepository.Update(fromAccount); err != nil {
		return err
	}

	if _, err := s.BankInfoRepository.Update(toAccount); err != nil {
		return err
	}

	return s.createTransaction(fromAccountNumber, toAccountNumber, amount, "transfer")
}

// Helper function to process a deposit transaction
func (s *TransactionValidator) processDeposit(toAccountNumber string, amount float64) error {
	toAccount, err := s.BankInfoRepository.GetByAccountNumber(toAccountNumber)
	if err != nil {
		return err
	}

	toAccount.Balance += amount

	if _, err := s.BankInfoRepository.Update(toAccount); err != nil {
		return err
	}

	return s.createTransaction("", toAccountNumber, amount, "deposit")
}

// Helper function to process a withdraw transaction
func (s *TransactionValidator) processWithdraw(fromAccountNumber string, amount float64) error {
	fromAccount, err := s.BankInfoRepository.GetByAccountNumber(fromAccountNumber)
	if err != nil {
		return err
	}

	if fromAccount.Balance < amount {
		return errors.New("insufficient balance")
	}

	fromAccount.Balance -= amount

	if _, err := s.BankInfoRepository.Update(fromAccount); err != nil {
		return err
	}

	return s.createTransaction(fromAccount.AccountNumber, "", amount, "withdraw")
}

// helper function to create a transaction record
func (s *TransactionValidator) createTransaction(fromAccountNumber, toAccountNumber string, amount float64, transactionType string) error {
	transaction := domain.Transaction{
		FromAccountNumber: fromAccountNumber,
		ToAccountNumber:   toAccountNumber,
		Amount:            amount,
		TransactionType:   transactionType,
	}

	if _, err := s.TransactionRepository.Create(&transaction); err != nil {
		return err
	}

	return nil
}
