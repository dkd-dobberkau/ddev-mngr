package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ddev-mngr",
	Short: "Manage DDEV projects",
	Long:  "A CLI tool to manage DDEV projects with interactive and command-line modes.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
