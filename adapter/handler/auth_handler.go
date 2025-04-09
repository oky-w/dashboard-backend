package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okyws/dashboard-backend/constants"
	"github.com/okyws/dashboard-backend/dto"
	"github.com/okyws/dashboard-backend/ports"
	"github.com/okyws/dashboard-backend/utils"
	"github.com/rs/zerolog/log"
)

// AuthHandler is the HTTP handler for the authentication service
type AuthHandler struct {
	Service ports.AuthService
}

// NewAuthHandler initializes a new AuthHandler instance.
func NewAuthHandler(service ports.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// Login handles the login route for authentication.
func (h *AuthHandler) Login(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	log.Info().Str("method", c.Request.Method).Str("path", c.Request.URL.Path).Msg("Initializing Login")

	var req dto.UserLoginDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.Service.LoginAccount(c, req.Username, req.Password)
	if err != nil {
		log.Error().Err(err).Msg("Username or password is incorrect. Failed to login")
		utils.ErrorResponse(c, http.StatusUnauthorized, "Username or password is incorrect. Failed to login")

		return
	}

	resp := dto.UserLoginResponseDTO(*data)

	utils.ResponseJSON(c, resp, http.StatusOK, "Login successful")
	log.Info().Str("username", req.Username).Msg("Login successful")
}
