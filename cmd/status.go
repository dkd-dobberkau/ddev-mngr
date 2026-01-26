package cmd

import (
	"fmt"
	"os"

	"github.com/dkd-dobberkau/ddev-mngr/internal/ddev"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status [project]",
	Short: "Show status of DDEV projects",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := ddev.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if len(args) == 1 {
			name := args[0]
			for _, p := range projects {
				if p.Name == name {
					fmt.Printf("%s: %s", p.Name, p.Status)
					if p.HTTPSUrl != "" && p.Status == "running" {
						fmt.Printf(" (%s)", p.HTTPSUrl)
					}
					fmt.Println()
					return
				}
			}
			fmt.Fprintf(os.Stderr, "Project '%s' not found\n", name)
			os.Exit(2)
		}

		// Show all
		running := 0
		stopped := 0
		for _, p := range projects {
			if p.Status == "running" {
				running++
			} else {
				stopped++
			}
		}
		fmt.Printf("Total: %d projects (%d running, %d stopped)\n", len(projects), running, stopped)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
