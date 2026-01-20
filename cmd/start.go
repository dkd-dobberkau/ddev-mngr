package cmd

import (
	"fmt"
	"os"

	"github.com/dkd-dobberkau/ddev-mngr/internal/ddev"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start <project>",
	Short: "Start a DDEV project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Starting %s... ", name)
		if err := ddev.Start(name); err != nil {
			fmt.Println("failed")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
