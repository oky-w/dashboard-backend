package services

import (
	"errors"

	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/ports"
	"gorm.io/gorm"
)

// CustomerService is the implementation of the customer service
type CustomerService struct {
	CustomerRepository ports.CustomerRepository
	UserRepository     ports.UserRepository // To validate UserID before creating a customer
}

// NewCustomerService creates a new customer service via dependency injection
func NewCustomerService(customerRepo ports.CustomerRepository, userRepo ports.UserRepository) *CustomerService {
	return &CustomerService{
		CustomerRepository: customerRepo,
		UserRepository:     userRepo,
	}
}

// CreateCustomer inserts a new customer into the database with UserID validation
func (s *CustomerService) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	existingUser, err := s.UserRepository.GetByID(customer.UserID.String())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	if existingUser == nil {
		return nil, errors.New("user ID does not exist")
	}

	existingCustomer, err := s.CustomerRepository.GetCustomerByUserID(customer.UserID.String())
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingCustomer != nil {
		return nil, errors.New("user already has a customer profile")
	}

	return s.CustomerRepository.Create(customer)
}

// GetCustomerByID fetches a customer by ID
func (s *CustomerService) GetCustomerByID(id string) (*domain.Customer, error) {
	return s.CustomerRepository.GetByID(id)
}

// GetCustomerByUserID fetches a customer by UserID
func (s *CustomerService) GetCustomerByUserID(userID string) (*domain.Customer, error) {
	return s.CustomerRepository.GetCustomerByUserID(userID)
}

// UpdateCustomer updates an existing customer
func (s *CustomerService) UpdateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	if customer.ID.String() == "" {
		return nil, errors.New("customer ID is required")
	}

	existingCustomer, err := s.CustomerRepository.GetByID(customer.ID.String())
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingCustomer == nil {
		return nil, errors.New("customer not found")
	}

	return s.CustomerRepository.Update(customer)
}

// DeleteCustomer removes a customer
func (s *CustomerService) DeleteCustomer(id string) error {
	return s.CustomerRepository.Delete(id)
}

// GetAllCustomers fetches all customers with pagination
func (s *CustomerService) GetAllCustomers(limit, offset int) ([]domain.Customer, error) {
	return s.CustomerRepository.GetAll(limit, offset)
}
