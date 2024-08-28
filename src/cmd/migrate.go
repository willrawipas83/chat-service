package cmd

import (
	"github.com/Rawipass/chat-service/logger"
	"github.com/Rawipass/chat-service/migration"
	"github.com/spf13/cobra"
)

var forceMigrate bool = false

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Chat Database",
	Run: func(cmd *cobra.Command, args []string) {
		// Init Logger
		logger.InitLogger()

		initComponent()

		initTimezone()

		migration.Migrate(false, -1, forceMigrate, false)

		logger.SyncLogger()
	},
}

func init() {
	rootCmd.AddCommand(MigrateCmd)
	MigrateCmd.PersistentFlags().BoolVar(&forceMigrate, "force", false, "force migrate (default is false)")
}
