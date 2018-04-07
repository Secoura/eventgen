package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var ConfigPath string

var EventGen = &cobra.Command{
	Use: "eventgen",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	addCommands()

	if c, err := EventGen.ExecuteC(); err != nil {
		c.Println("")
		os.Exit(-1)
	}
}

func addCommands() {
	EventGen.AddCommand(versionCmd)
}

func init() {
	EventGen.PersistentFlags().StringVarP(&ConfigPath, "configPath", "c", "config.yaml", "the path to the app config file")
}
