package config

import (
	"github.com/okyws/dashboard-backend/constants"
	"github.com/okyws/dashboard-backend/domain"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

// NewRedisClient creates a new Redis client
func NewRedisClient(config *domain.Configuration) *redis.Client {
	redisClient := redis.NewClient(config.GetRedisConfig())

	log.Info().Msg(constants.MsgRedisConnectSuccess)

	return redisClient
}

// CloseRedisClient closes the Redis client
func CloseRedisClient(redisClient *redis.Client) {
	err := redisClient.Close()
	if err != nil {
		log.Error().Err(err).Msg(constants.MsgRedisCloseFail)
	}

	log.Info().Msg(constants.MsgRedisCloseSuccess)
}
