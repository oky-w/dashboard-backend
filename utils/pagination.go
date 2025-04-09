package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPaginationParams extracts "limit" and "offset" from query parameters
func GetPaginationParams(c *gin.Context) (int, int) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	// Default values
	limit := 10
	offset := 0

	// Convert "limit" from string to int
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	// Convert "offset" from string to int
	if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
		offset = o
	}

	return limit, offset
}
