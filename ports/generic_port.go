// Package ports contains the interfaces for repositories and services
package ports

// GenericRepository is a generic interface for repositories
type GenericRepository[T any] interface {
	Create(entity *T) (*T, error)
	GetByID(id string) (*T, error)
	Update(entity *T) (*T, error)
	Delete(id string) error
	GetAll(limit, offset int) ([]T, error)
}

// GenericService is a generic interface for services
type GenericService[T any] interface {
	Create(entity *T) (*T, error)
	GetByID(id string) (*T, error)
	Update(entity *T) (*T, error)
	Delete(id string) error
	GetAll(limit, offset int) ([]T, error)
}
