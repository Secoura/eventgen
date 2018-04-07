package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Long:  "Prints the version and build time.",
	RunE: func(cmd *cobra.Command, args []string) error {
		printVersion()
		return nil
	},
}

func printVersion() {
	fmt.Println("Event Gen - v0.0.1")
}
