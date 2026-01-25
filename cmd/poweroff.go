package cmd

import (
	"fmt"
	"os"

	"github.com/dkd-dobberkau/ddev-mngr/internal/ddev"
	"github.com/spf13/cobra"
)

var poweroffCmd = &cobra.Command{
	Use:   "poweroff",
	Short: "Stop all DDEV projects and the router",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Stopping all DDEV projects... ")
		if err := ddev.Poweroff(); err != nil {
			fmt.Println("failed")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(poweroffCmd)
}
