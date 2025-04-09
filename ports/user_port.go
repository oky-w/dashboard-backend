// Package ports contains the interfaces for repositories and services
package ports

import (
	"github.com/okyws/dashboard-backend/domain"
)

// UserRepository is the interface for the user repository
type UserRepository interface {
	GenericRepository[domain.User]
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
}

// UserService is the interface for the user service
type UserService interface {
	GenericService[domain.User]
	GetUserByUsername(username string) (*domain.User, error)
}
