package ports

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/okyws/dashboard-backend/dto"
)

// AuthRepository is the interface for the authentication repository
type AuthRepository interface {
	SaveToken(ctx *gin.Context, userID uuid.UUID, token, expiresAt string) error
	GetTokenExpiration(ctx *gin.Context, token string) (*time.Time, error)
}

// AuthService is the interface for the authentication service
type AuthService interface {
	LoginAccount(ctx *gin.Context, username, password string) (*dto.UserLoginResponseDTO, error)
	ValidateToken(ctx *gin.Context, token string) (bool, error)
}
