// Package domain contains the application configuration settings.
package domain

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/okyws/dashboard-backend/constants"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

// Configuration struct for application configuration
type Configuration struct {
	AppName    string
	AppVersion string
	ServerPort string
	ServerHost string
	APIKey     string
	JWTSecret  string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBSSLMode  string
	DBTimeZone string
	REDISHost  string
	REDISPort  string
	REDISDB    string
	REDISPass  string
}

// LoadConfig reads configuration values from .env
func LoadConfig() (*Configuration, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg(constants.MsgConfigLoadFail)
	}

	config := &Configuration{
		AppName:    os.Getenv("APP_NAME"),
		AppVersion: os.Getenv("APP_VERSION"),
		ServerPort: os.Getenv("SERVER_PORT"),
		ServerHost: os.Getenv("SERVER_HOST"),
		APIKey:     os.Getenv("API_KEY"),
		JWTSecret:  os.Getenv("JWT_SECRET_KEY"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSL_MODE"),
		DBTimeZone: os.Getenv("DB_TIMEZONE"),
		REDISHost:  os.Getenv("REDIS_HOST"),
		REDISPort:  os.Getenv("REDIS_PORT"),
		REDISDB:    os.Getenv("REDIS_DB"),
		REDISPass:  os.Getenv("REDIS_PASS"),
	}

	return config, nil
}

// GetDatabaseConfig returns the database source name
func (c *Configuration) GetDatabaseConfig() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, c.DBSSLMode, c.DBTimeZone,
	)
}

// GetRedisConfig returns the redis configuration
func (c *Configuration) GetRedisConfig() *redis.Options {
	db, _ := strconv.Atoi(c.REDISDB)

	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.REDISHost, c.REDISPort),
		Password: c.REDISPass,
		DB:       db,
	}
}
