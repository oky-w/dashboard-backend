// Package utils contains utility functions
package utils

import (
	"errors"
	"time"
)

// FormatDate converts a date string to a time.Time object
func FormatDate(date string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("invalid date format, should be YYYY-MM-DD")
	}

	return &t, nil
}
