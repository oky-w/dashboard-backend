// Package utils contains utility functions for sending JSON responses
package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/okyws/dashboard-backend/dto"
	"github.com/rs/zerolog/log"
)

// ResponseJSON sends a successful JSON response in Gin
func ResponseJSON(c *gin.Context, data interface{}, code int, message string) {
	resp := dto.SuccessResponseDTO[interface{}]{
		Status:  "success",
		Code:    code,
		Message: message,
		Data:    data,
	}

	sendJSON(c, resp, code, message)
}

// ErrorResponse sends an error JSON response in Gin
func ErrorResponse(c *gin.Context, code int, message string) {
	resp := dto.ErrorResponseDTO{
		Status:  "error",
		Code:    code,
		Message: message,
	}

	sendJSON(c, resp, code, message)
}

// sendJSON is a helper function for sending JSON responses
func sendJSON(c *gin.Context, resp interface{}, code int, message string) {
	c.JSON(code, resp)

	// Log response
	log.Info().
		Int("code", code).
		Str("message", message).
		Msg("Response sent successfully")
}
