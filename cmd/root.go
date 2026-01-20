package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dkd-dobberkau/ddev-mngr/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ddev-mngr",
	Short: "Manage DDEV projects",
	Long:  "A CLI tool to manage DDEV projects with interactive and command-line modes.",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.NewModel())
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
