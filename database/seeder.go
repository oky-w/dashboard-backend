// Package database contains the database migration and seeding logic.
package database

import (
	"crypto/rand"
	"encoding/binary"
	"strings"

	"github.com/google/uuid"
	"github.com/okyws/dashboard-backend/constants"
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/utils"
	"github.com/rs/zerolog/log"
)

// UserSeed returns a list of users for seeding
func UserSeed() []domain.User {
	hash := "$2a$10$AR6SC2Fh1OHJM9SGH1CsWOgm5GwiuvwKq3GvdtDvXNCqwHiUNkt4e"

	users := []domain.User{
		{
			ID:       uuid.New(),
			Email:    "user@example.com",
			Username: "user",
			Password: hash,
			Role:     "user",
		},
		{
			ID:       uuid.New(),
			Email:    "john.doe1@example.com",
			Username: "john.doe1",
			Password: hash,
			Role:     "user",
		},
		{
			ID:       uuid.New(),
			Email:    "jane.doe2@example.com",
			Username: "jane.doe2",
			Password: hash,
			Role:     "user",
		},
		{
			ID:       uuid.New(),
			Email:    "admin@example.com",
			Username: "admin",
			Password: hash,
			Role:     "admin",
		},
	}

	return users
}

// CustomerSeed returns a list of customers for seeding
func CustomerSeed(users []domain.User) []domain.Customer {
	date, _ := utils.FormatDate("1990-01-01")
	customers := make([]domain.Customer, len(users))

	for i, user := range users {
		phoneNumber, _ := utils.GenerateAccountNumber(10)

		customers[i] = domain.Customer{
			UserID:      user.ID,
			FullName:    strings.Split(user.Username, ".")[0],
			PhoneNumber: phoneNumber,
			DateOfBirth: *date,
			Address:     "123 Main Street",
		}
	}

	return customers
}

// BankAccountSeed returns a list of bank accounts for seeding
func BankAccountSeed(users []domain.User) []domain.BankAccount {
	bankAccounts := make([]domain.BankAccount, 0)

	for _, user := range users {
		accountNumber, _ := utils.GenerateAccountNumber(10)

		// Main account
		bankAccounts = append(bankAccounts, domain.BankAccount{
			UserID:        user.ID,
			AccountType:   "rekening-utama",
			AccountNumber: accountNumber,
			Balance:       500000,
			AccountStatus: true,
		})

		// Saku accounts (Max 8)
		for i := 0; i < 8; i++ {
			accountNumber, _ := utils.GenerateAccountNumber(10)
			bankAccounts = append(bankAccounts, domain.BankAccount{
				UserID:        user.ID,
				AccountType:   "saku",
				AccountNumber: accountNumber,
				Balance:       0,
				AccountStatus: true,
			})
		}

		// Deposit accounts (Max 3)
		for i := 0; i < 3; i++ {
			accountNumber, _ := utils.GenerateAccountNumber(10)
			bankAccounts = append(bankAccounts, domain.BankAccount{
				UserID:        user.ID,
				AccountType:   "deposito",
				AccountNumber: accountNumber,
				Balance:       0,
				AccountStatus: true,
			})
		}
	}

	return bankAccounts
}

// TransactionSeed returns a list of transactions for seeding
func TransactionSeed(users []domain.User, bankAccounts []domain.BankAccount) []domain.Transaction {
	var transactions []domain.Transaction

	// Ensure that we have enough bank accounts
	accountMap := make(map[uuid.UUID][]domain.BankAccount)
	for _, account := range bankAccounts {
		accountMap[account.UserID] = append(accountMap[account.UserID], account)
	}

	// Create random transactions for each user
	for _, user := range users {
		accounts := accountMap[user.ID]

		// Create 5 random transactions per user
		for i := 0; i < 5; i++ {
			transactionType := randomTransactionType()

			// Generate a secure random amount between 100 and 10000
			amount := secureRandomFloat(100, 10000)

			var transaction domain.Transaction

			// Pick a random index for accounts
			fromAccount := accounts[secureRandomInt(len(accounts))]
			toAccount := accounts[secureRandomInt(len(accounts))]

			// Ensure that fromAccount and toAccount are different for 'transfer'
			if transactionType == "transfer" && fromAccount.AccountNumber == toAccount.AccountNumber {
				// Re-pick toAccount if it is the same as fromAccount
				toAccount = accounts[secureRandomInt(len(accounts))]
			}

			switch transactionType {
			case "transfer":
				// Create a transfer transaction
				transaction = domain.Transaction{
					ID:                uuid.New(),
					FromAccountNumber: fromAccount.AccountNumber,
					ToAccountNumber:   toAccount.AccountNumber,
					Amount:            amount,
					TransactionType:   "transfer",
					Status:            "success",
				}

			case "deposit":
				// Pick a random account for deposit
				transaction = domain.Transaction{
					ID:                uuid.New(),
					FromAccountNumber: "",
					ToAccountNumber:   toAccount.AccountNumber,
					Amount:            amount,
					TransactionType:   "deposit",
					Status:            "success",
				}

			case "withdraw":
				// Pick a random account for withdraw
				transaction = domain.Transaction{
					ID:                uuid.New(),
					FromAccountNumber: fromAccount.AccountNumber,
					ToAccountNumber:   "",
					Amount:            amount,
					TransactionType:   "withdraw",
					Status:            "success",
				}
			}

			// Append the generated transaction to the list
			transactions = append(transactions, transaction)
		}
	}

	return transactions
}

// randomTransactionType returns a random transaction type
func randomTransactionType() string {
	transactionTypes := []string{"transfer", "deposit", "withdraw"}
	return transactionTypes[secureRandomInt(len(transactionTypes))]
}

// secureRandomInt generates a secure random int in the range [0, limit)
func secureRandomInt(limit int) int {
	b := make([]byte, 1)
	_, err := rand.Read(b)

	if err != nil {
		log.Fatal().Err(err).Msg(constants.MsgRandomNumberError)
	}

	return int(b[0]) % limit
}

// secureRandomFloat generates a secure random float in the range [min, max)
func secureRandomFloat(lowerLimit, highLimit float64) float64 {
	randomByte := make([]byte, 8)
	_, err := rand.Read(randomByte)

	if err != nil {
		log.Fatal().Err(err).Msg(constants.MsgRandomNumberError)
	}

	randomInt := binary.LittleEndian.Uint64(randomByte)

	return lowerLimit + float64(randomInt%(uint64(highLimit)-uint64(lowerLimit)))
}
