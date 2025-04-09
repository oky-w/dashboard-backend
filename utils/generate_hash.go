// Package utils contains utility functions for generating bcrypt hashes
package utils

import "golang.org/x/crypto/bcrypt"

// GeneratePasswordHash generates a bcrypt hash from the given password
func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
