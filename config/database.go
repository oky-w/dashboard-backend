// Package config contains the application configuration settings.
package config

import (
	"time"

	"github.com/okyws/dashboard-backend/constants"
	"github.com/okyws/dashboard-backend/database"
	"github.com/okyws/dashboard-backend/domain"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBConnectionENV initializes a new database connection
func NewDBConnectionENV(config *domain.Configuration) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDatabaseConfig()), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg(constants.MsgDBConnectFail)
		return nil, err
	}

	log.Info().Msg(constants.MsgDBConnectSuccess)

	return db, nil
}

// CloseDatabase closes the database connection
func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg(constants.MsgDBCloseFail)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal().Err(err).Msg(constants.MsgDBCloseFail)
	} else {
		log.Info().Msg(constants.MsgDBCloseSuccess)
	}
}

// GetDatabaseConnection gets the database connection
func GetDatabaseConnection() (*gorm.DB, *redis.Client, error) {
	configuration, err := domain.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Str("error", err.Error()).Msg(constants.MsgConfigLoadFail)
		return nil, nil, err
	}

	db, err := NewDBConnectionENV(configuration)
	if err != nil {
		log.Fatal().Err(err).Str("error", err.Error()).Msg(constants.MsgDBConnectFail)
		return nil, nil, err
	}

	redisClient := NewRedisClient(configuration)

	log.Info().Msg(constants.MsgDBConnectSuccess)

	return db, redisClient, nil
}

// MigrateDB migrates the database schema
func MigrateDB(db *gorm.DB) {
	if err := db.AutoMigrate(&domain.User{}, &domain.Customer{}, &domain.BankAccount{}, &domain.Transaction{}); err != nil {
		log.Fatal().Err(err).Msg(constants.MsgDBMigrateFail)
	}

	log.Info().Msg(constants.MsgDBMigrateSuccess)
}

// DropDB drops the database schema
func DropDB(db *gorm.DB) {
	if err := db.Migrator().DropTable(&domain.User{}, &domain.Customer{}, &domain.BankAccount{}, &domain.Transaction{}); err != nil {
		log.Fatal().Err(err).Msg(constants.MsgDBDropFail)
	}

	log.Info().Msg(constants.MsgDBDropSuccess)
}

// SeedDB seeds the database with initial data
func SeedDB(db *gorm.DB) {
	start := time.Now()
	users := database.UserSeed()
	db.CreateInBatches(users, 100)

	customers := database.CustomerSeed(users)
	db.CreateInBatches(customers, 100)

	bankAccounts := database.BankAccountSeed(users)
	db.CreateInBatches(bankAccounts, 100)

	transactions := database.TransactionSeed(users, bankAccounts)
	db.CreateInBatches(transactions, 100)

	duration := time.Since(start)
	log.Info().Str("duration", duration.String()).Msg(constants.MsgDBSeedSuccess)
}
