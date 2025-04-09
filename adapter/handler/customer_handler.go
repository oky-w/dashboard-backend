// Package handler contains the HTTP handlers for the customer service
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/okyws/dashboard-backend/constants"
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/dto"
	"github.com/okyws/dashboard-backend/services"
	"github.com/okyws/dashboard-backend/utils"
	"gorm.io/gorm"
)

// CustomerHandlerAdapter is the HTTP handler for the customer service
type CustomerHandlerAdapter struct {
	CustomerService *services.CustomerService
}

// NewCustomerHandler creates a new customer handler via dependency injection
func NewCustomerHandler(service *services.CustomerService) *CustomerHandlerAdapter {
	return &CustomerHandlerAdapter{CustomerService: service}
}

// HandleCreateCustomer handles the HTTP request for creating a new customer
func (h *CustomerHandlerAdapter) HandleCreateCustomer(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	var req dto.CustomerCreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	formattedDate, err := utils.FormatDate(req.DateOfBirth)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	customer, err := h.CustomerService.CreateCustomer(&domain.Customer{
		UserID:      req.UserID,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: *formattedDate,
		Address:     req.Address,
	})

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusUnprocessableEntity, constants.MsgUnprocessable)
		return
	}

	customerDTO := *domain.MapCustomerToDTO(customer)

	utils.ResponseJSON(c, customerDTO, http.StatusCreated, "Customer created successfully")
}

// HandleGetCustomerByID handles the HTTP request for getting a customer by ID
func (h *CustomerHandlerAdapter) HandleGetCustomerByID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	customer, err := h.CustomerService.GetCustomerByID(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if customer == nil || err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	customerDTO := *domain.MapCustomerToDTO(customer)

	utils.ResponseJSON(c, customerDTO, http.StatusOK, "Customer fetched successfully")
}

// HandleGetCustomerByUserID handles the HTTP request for getting a customer by UserID
func (h *CustomerHandlerAdapter) HandleGetCustomerByUserID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	userID := c.Param("user_id")

	customer, err := h.CustomerService.GetCustomerByUserID(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if customer == nil || err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	customerDTO := *domain.MapCustomerToDTO(customer)

	utils.ResponseJSON(c, customerDTO, http.StatusOK, "Customer fetched successfully")
}

// HandleUpdateCustomer handles the HTTP request for updating an existing customer
func (h *CustomerHandlerAdapter) HandleUpdateCustomer(c *gin.Context) {
	if c.Request.Method != http.MethodPut {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	var req dto.CustomerUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	formattedDate, err := utils.FormatDate(req.DateOfBirth)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	customer, err := h.CustomerService.UpdateCustomer(&domain.Customer{
		ID:          uuid.MustParse(id),
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: *formattedDate,
		Address:     req.Address,
	})

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound || customer == nil {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	customerDTO := *domain.MapCustomerToDTO(customer)

	utils.ResponseJSON(c, customerDTO, http.StatusOK, "Customer updated successfully")
}

// HandleDeleteCustomer handles the HTTP request for deleting a customer
func (h *CustomerHandlerAdapter) HandleDeleteCustomer(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	err := h.CustomerService.DeleteCustomer(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	utils.ResponseJSON(c, nil, http.StatusOK, "Customer deleted successfully")
}

// HandleGetAllCustomers handles the HTTP request for getting all customers with pagination
func (h *CustomerHandlerAdapter) HandleGetAllCustomers(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	limit, offset := utils.GetPaginationParams(c)

	customers, err := h.CustomerService.GetAllCustomers(limit, offset)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	customerDTOs := make([]dto.CustomerDTO, len(customers))
	for i, customer := range customers {
		customerDTOs[i] = *domain.MapCustomerToDTO(&customer)
	}

	utils.ResponseJSON(c, customerDTOs, http.StatusOK, "Customers fetched successfully")
}
