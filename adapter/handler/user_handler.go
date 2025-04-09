// Package handler contains the HTTP handlers for the user service
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

// UserHandlerAdapter is the HTTP handler for the user service
type UserHandlerAdapter struct {
	UserService *services.UserService
}

// NewUserHandler creates a new user handler via dependency injection
func NewUserHandler(service *services.UserService) *UserHandlerAdapter {
	return &UserHandlerAdapter{UserService: service}
}

// HandleCreateUser implements the HTTP handler for creating a new user
func (h *UserHandlerAdapter) HandleCreateUser(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	var req dto.UserCreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.UserService.CreateUser(&domain.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	})

	if err != nil {
		utils.ErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	userDTO := dto.UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}

	utils.ResponseJSON(c, userDTO, http.StatusCreated, "User created successfully")
}

// HandleGetUserByID implements the HTTP handler for getting a user by ID
func (h *UserHandlerAdapter) HandleGetUserByID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	user, err := h.UserService.GetUserByID(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	userDTO := dto.UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}

	utils.ResponseJSON(c, userDTO, http.StatusOK, "User fetched successfully")
}

// HandleGetUserByUsername implements the HTTP handler for getting a user by username
func (h *UserHandlerAdapter) HandleGetUserByUsername(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	username := c.Param("username")

	user, err := h.UserService.GetUserByUsername(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	userDTO := dto.UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}

	utils.ResponseJSON(c, userDTO, http.StatusOK, "User fetched successfully")
}

// HandleGetAllUsers implements the HTTP handler for getting all users with pagination
func (h *UserHandlerAdapter) HandleGetAllUsers(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	limit, offset := utils.GetPaginationParams(c)

	users, err := h.UserService.GetAllUsers(limit, offset)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserDTO{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Role:     user.Role,
		}
	}

	utils.ResponseJSON(c, userDTOs, http.StatusOK, "Users fetched successfully")
}

// HandleUpdateUser implements the HTTP handler for updating a user
func (h *UserHandlerAdapter) HandleUpdateUser(c *gin.Context) {
	if c.Request.Method != http.MethodPut {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	var req dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.UserService.UpdateUser(&domain.User{
		ID:       uuid.MustParse(id),
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	})

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	userDTO := dto.UserDTO{
		ID:       uuid.MustParse(id),
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}

	utils.ResponseJSON(c, userDTO, http.StatusOK, "User updated successfully")
}

// HandleDeleteUser implements the HTTP handler for deleting a user
func (h *UserHandlerAdapter) HandleDeleteUser(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	err := h.UserService.DeleteUser(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	utils.ResponseJSON(c, nil, http.StatusOK, "User deleted successfully")
}
