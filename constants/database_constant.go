// Package constants contains constants for the application
package constants

// Constants for Log Database Messages
const (
	MsgDBConnectSuccess    = "Successfully connected to database"
	MsgDBConnectFail       = "Failed to connect to database"
	MsgDBCloseFail         = "Failed to close database connection"
	MsgDBCloseSuccess      = "Database connection closed successfully"
	MsgDBMigrateSuccess    = "Successfully migrated database"
	MsgDBMigrateFail       = "Failed to migrate database"
	MsgDBDropSuccess       = "Successfully dropped database"
	MsgDBDropFail          = "Failed to drop database"
	MsgDBSeedSuccess       = "Successfully seeded database"
	MsgRedisConnectFail    = "Failed to connect to redis"
	MsgRedisConnectSuccess = "Successfully connected to redis"
	MsgRedisCloseFail      = "Failed to close redis connection"
	MsgRedisCloseSuccess   = "Redis connection closed successfully"
)
