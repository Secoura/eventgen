package commands

import (
	"io/ioutil"
	"log"

	"github.com/secoura/eventgen/config"
	"github.com/secoura/eventgen/processor"
	"github.com/spf13/cobra"
	"errors"
	"fmt"
	"strings"
)

var templateDir = "templates"

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate events",
	Long:  "Generate events using the provided template",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generateEvents()
	},
}

func generateEvents() error {
	data, err := ioutil.ReadFile(templateDir + "/" + config.GetConfig().Template)
	if err != nil {
		log.Fatal("Error reading template file: ", err.Error())
		return errors.New("failed to read template file")
	}
	template := strings.TrimSpace(string(data))

	delimiter := "\n"
	if d := config.GetConfig().Delimiter; d != "" {
		delimiter = d
	}

	var i int
	var result string
	for i = 0; i < config.GetConfig().NumberOfEvents; i++ {
		result = processor.ProcessTemplate(template)
		fmt.Printf("%s%s", result, delimiter)
	}
	return nil
}
