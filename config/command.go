package config

import (
	"os"

	"github.com/okyws/dashboard-backend/constants"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// DropCmd command to drop the database
var DropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop the database",
	Run: func(_ *cobra.Command, _ []string) {
		log.Info().Msg("Dropping the database...")

		db, _, err := GetDatabaseConnection()
		if err != nil {
			return
		}
		defer CloseDatabase(db)

		DropDB(db)
		log.Info().Msg("Database dropped successfully.")
	},
}

// SeedCmd command to seed the database
var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database",
	Run: func(_ *cobra.Command, _ []string) {
		log.Info().Msg("Seeding the database...")

		db, _, err := GetDatabaseConnection()
		if err != nil {
			return
		}
		defer CloseDatabase(db)

		SeedDB(db)
		log.Info().Msg("Database seeded successfully.")
	},
}

// MigrateCmd command to migrate the database
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(_ *cobra.Command, _ []string) {
		log.Info().Msg("Migrating the database...")

		db, _, err := GetDatabaseConnection()
		if err != nil {
			return
		}
		defer CloseDatabase(db)

		MigrateDB(db)
		log.Info().Msg("Database migrated successfully.")
	},
}

// InitCommand initializes the command
func InitCommand() {
	var rootCmd = &cobra.Command{Use: "dbtool"}

	rootCmd.AddCommand(SeedCmd, DropCmd, MigrateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg(constants.MsgCommandFail)
		os.Exit(1)
	}
}
