// Package utils contains utility functions for validating bcrypt hashes
package utils

import "regexp"

// IsValidBcryptHash checks if the given string is a valid bcrypt hash
func IsValidBcryptHash(hash string) bool {
	regex := `^\$2[aby]\$\d{2}\$[./A-Za-z0-9]{53}$`
	re := regexp.MustCompile(regex)

	return re.MatchString(hash)
}
