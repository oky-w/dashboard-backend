package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okyws/dashboard-backend/config"
	"github.com/okyws/dashboard-backend/dto"
	"github.com/okyws/dashboard-backend/ports"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// AuthAdapter is the implementation of the authentication service
type AuthAdapter struct {
	repo ports.AuthRepository
	user ports.UserRepository
}

// NewAuthService creates a new authentication service
func NewAuthService(repo ports.AuthRepository, user ports.UserRepository) *AuthAdapter {
	return &AuthAdapter{repo: repo, user: user}
}

// LoginAccount logs in a user
func (u *AuthAdapter) LoginAccount(ctx *gin.Context, username, password string) (*dto.UserLoginResponseDTO, error) {
	log.Info().Str("username", username).Msg("LoginAccount started")

	user, err := u.user.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Error().Err(err).Msg("Username or password is incorrect. Failed to login")
		return nil, errors.New("username or password is incorrect")
	}

	token, expiresAt, err := config.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate token")
		return nil, errors.New("could not generate token")
	}

	err = u.repo.SaveToken(ctx, user.ID, token, expiresAt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to save token")
		return nil, fmt.Errorf("failed to save token: %v", err)
	}

	log.Info().Str("username", user.Username).Str("expiresAt", expiresAt).Msg("Login success")

	return &dto.UserLoginResponseDTO{Username: user.Username, UserID: user.ID.String(), Role: user.Role, Token: token, ExpiresAt: expiresAt}, nil
}

// ValidateToken validates a token and returns true if it is valid
func (u *AuthAdapter) ValidateToken(ctx *gin.Context, token string) (bool, error) {
	expiresAt, err := u.repo.GetTokenExpiration(ctx, token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate token")
		return false, fmt.Errorf("failed to validate token: %v", err)
	}

	// Check if the token is expired
	if expiresAt.Before(time.Now()) {
		log.Info().Str("token", token).Msg("Token is expired")
		return false, nil
	}

	log.Info().Msg("Token is valid")

	return true, nil
}
