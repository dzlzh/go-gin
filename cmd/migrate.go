package cmd

import (
	"go-gin/initialize"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "执行迁移",
	Long:  "执行Mysql迁移",
	Run:   migrateRun,
}

func migrateRun(cmd *cobra.Command, args []string) {
	initialize.AutoMigrate()
}
