package config

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Claims struct for JWT
type Claims struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}

// Secret key for JWT
var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateJWT for generating JWT token
func GenerateJWT(id uuid.UUID, username, role string) (string, string, error) {
	log.Info().Msg("Initializing Generate JWT Token")

	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		ID:       id,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	converedtTime := claims.ExpiresAt.Format("2006-01-02 15:04:05")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate token")
		return "", "", err
	}

	log.Info().Str("username", username).Str("expiresAt", converedtTime).Msg("Token generated successfully")

	return tokenString, converedtTime, nil
}

// ParseToken for validating JWT
func ParseToken(tokenString string) (*Claims, error) {
	log.Info().Msg("Initializing Parse JWT Token")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to parse token")
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		log.Error().Msg("Invalid token")
		return nil, errors.New("invalid token")
	}

	log.Info().Str("username", claims.Username).Str("expiresAt", claims.ExpiresAt.String()).Msg("Token parsed successfully")

	return claims, nil
}
