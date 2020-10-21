package cmd

import (
	"go-gin/cmd/cron"

	"github.com/spf13/cobra"
)

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "定时任务",
	Long:  "定时任务",
	Run:   func(cmd *cobra.Command, args []string) { cron.Cron() },
}
