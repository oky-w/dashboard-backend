// Package repository contains the adapters for the customer repository
package repository

import (
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
	"gorm.io/gorm"
)

// CustomerRepositoryAdapter is the adapter for the customer repository
type CustomerRepositoryAdapter struct {
	db *gorm.DB
}

// NewCustomerRepositoryAdapter creates a new customer repository adapter via dependency injection
func NewCustomerRepositoryAdapter(db *gorm.DB) ports.CustomerRepository {
	return &CustomerRepositoryAdapter{db: db}
}

// Create inserts a new customer into the database using a transaction.
func (r *CustomerRepositoryAdapter) Create(customer *domain.Customer) (*domain.Customer, error) {
	var createdCustomer domain.Customer

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(customer).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Preload("User").First(&createdCustomer, "id = ?", customer.ID).Error
	})

	if err != nil {
		return nil, err
	}

	return &createdCustomer, nil
}

// GetByID fetches a customer by ID, returning nil if not found.
func (r *CustomerRepositoryAdapter) GetByID(id string) (*domain.Customer, error) {
	var customer domain.Customer

	err := r.db.Preload("User").First(&customer, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &customer, nil
}

// GetCustomerByUserID fetches a customer by associated User ID.
func (r *CustomerRepositoryAdapter) GetCustomerByUserID(userID string) (*domain.Customer, error) {
	var customer domain.Customer

	err := r.db.Preload("User").First(&customer, "user_id = ?", userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &customer, nil
}

// GetCustomerByPhoneNumber fetches a customer by phone number.
func (r *CustomerRepositoryAdapter) GetCustomerByPhoneNumber(phoneNumber string) (*domain.Customer, error) {
	var customer domain.Customer

	err := r.db.Preload("User").First(&customer, "phone_number = ?", phoneNumber).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &customer, nil
}

// Update updates an existing customer in the database.
func (r *CustomerRepositoryAdapter) Update(customer *domain.Customer) (*domain.Customer, error) {
	var updatedCustomer domain.Customer

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("User").First(&updatedCustomer, "id = ?", customer.ID).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Preload("User").Model(&updatedCustomer).Updates(customer).Error
	})

	if err != nil {
		return nil, err
	}

	return &updatedCustomer, nil
}

// Delete removes a customer by ID and ensures that a customer was actually deleted.
func (r *CustomerRepositoryAdapter) Delete(id string) error {
	result := r.db.Delete(&domain.Customer{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAll fetches customers with pagination.
func (r *CustomerRepositoryAdapter) GetAll(limit, offset int) ([]domain.Customer, error) {
	var customers []domain.Customer
	err := r.db.Preload("User").Limit(limit).Offset(offset).Find(&customers).Error

	return customers, err
}
