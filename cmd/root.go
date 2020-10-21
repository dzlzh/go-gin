package cmd

import (
	"go-gin/bootstrap"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) { bootstrap.Run() },
}

func init() {
	rootCmd.AddCommand(cronCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
