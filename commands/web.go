package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/yeoji/eventgen/config"
	"github.com/yeoji/eventgen/processor"
)

var webTemplateDir = "templates/web"

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Generate web events",
	Long:  "Generate web events in the provided log template",
	RunE: func(cmd *cobra.Command, args []string) error {
		generateWebEvents()
		return nil
	},
}

func generateWebEvents() {
	data, _ := ioutil.ReadFile(webTemplateDir + "/" + Template)
	template := string(data)

	var i int
	var result string
	for i = 0; i < config.GetConfig().NoOfEvents; i++ {
		result = processor.ProcessTemplate(template)
		fmt.Println(result)
	}
}
