package cmd

import (
	"fmt"
	"os"

	"github.com/oliverguenther/ddev-mngr/internal/ddev"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all DDEV projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := ddev.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		for _, p := range projects {
			symbol := "○"
			if p.Status == "running" {
				symbol = "●"
			}
			fmt.Printf("%s %-25s %s\n", symbol, p.Name, p.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
