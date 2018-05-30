package commands

import (
	"os"

	"github.com/secoura/eventgen/config"
	"github.com/secoura/eventgen/processor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	EventGen.AddCommand(generateCmd)
}

func init() {
	cobra.OnInitialize(config.LoadConfig)
	viper.BindPFlags(EventGen.PersistentFlags())
	processor.RegisterProcessors()
}
