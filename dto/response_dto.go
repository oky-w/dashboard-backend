// Package dto contains the data transfer objects for the application
package dto

// SuccessResponseDTO is the response structure for successful requests
type SuccessResponseDTO[T any] struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// ErrorResponseDTO is the response structure for failed requests
type ErrorResponseDTO struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
