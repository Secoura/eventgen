package commands

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Secoura/eventgen/config"
	"github.com/Secoura/eventgen/processor"
	"github.com/spf13/cobra"
)

var templateDir = "templates"

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate events",
	Long:  "Generate events using the provided template",
	RunE: func(cmd *cobra.Command, args []string) error {
		generateEvents()
		return nil
	},
}

func generateEvents() {
	data, err := ioutil.ReadFile(templateDir + "/" + config.GetConfig().Template)
	if err != nil {
		log.Fatal("Error reading template file: ", err.Error())
	}
	template := string(data)

	var i int
	var result string
	for i = 0; i < config.GetConfig().NoOfEvents; i++ {
		result = processor.ProcessTemplate(template)
		fmt.Println(result)
	}
}
