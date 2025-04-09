// Package ports contains the interfaces for repositories and services
package ports

import "github.com/okyws/dashboard-backend/domain"

// CustomerRepository is the interface for the customer repository
type CustomerRepository interface {
	GenericRepository[domain.Customer]
	GetCustomerByUserID(userID string) (*domain.Customer, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*domain.Customer, error)
}

// CustomerService is the interface for the customer service
type CustomerService interface {
	GenericService[domain.Customer]
	GetCustomerByUserID(userID string) (*domain.Customer, error)
}
