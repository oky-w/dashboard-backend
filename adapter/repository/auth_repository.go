package repository

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

// AuthRepositoryRedis is the implementation of the authentication repository using Redis
type AuthRepositoryRedis struct {
	RedisClient *redis.Client
}

// NewAuthRepositoryRedis creates a new instance of AuthRepositoryRedis
func NewAuthRepositoryRedis(redisClient *redis.Client) *AuthRepositoryRedis {
	return &AuthRepositoryRedis{RedisClient: redisClient}
}

// SaveToken stores the token and its expiration time in Redis
func (r *AuthRepositoryRedis) SaveToken(ctx *gin.Context, userID uuid.UUID, token, expiresAt string) error {
	expirationTime, err := time.Parse("2006-01-02 15:04:05", expiresAt)
	if err != nil {
		return fmt.Errorf("failed to parse expires_at: %v", err)
	}

	err = r.RedisClient.SetEx(ctx, token, userID, time.Until(expirationTime)).Err()
	if err != nil {
		return fmt.Errorf("failed to save token to Redis: %v", err)
	}

	log.Info().Str("userID", userID.String()).Msg("Token saved to Redis")

	return nil
}

// GetTokenExpiration retrieves the expiration time of a token from Redis
func (r *AuthRepositoryRedis) GetTokenExpiration(ctx *gin.Context, token string) (*time.Time, error) {
	expiresAtStr, err := r.RedisClient.HGet(ctx, token, "expires_at").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token expiration from Redis: %v", err)
	}

	// Parse the expiration time string into a time.Time object
	expiresAt, err := time.Parse("2006-01-02 15:04:05", expiresAtStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expires_at: %v", err)
	}

	log.Info().Msg("Token expiration retrieved from Redis")

	return &expiresAt, nil
}
