// Package utils contains utility functions
package utils

import "crypto/rand"

// GenerateAccountNumber generates a random account number
func GenerateAccountNumber(length int) (string, error) {
	letters := "0123456789"
	accountNumber := make([]byte, length)

	for i := range accountNumber {
		buf := make([]byte, 1)

		_, err := rand.Read(buf)
		if err != nil {
			return "", err
		}

		accountNumber[i] = letters[buf[0]%byte(len(letters))]
	}

	return string(accountNumber), nil
}
