package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
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
	for i = 0; i < 1000; i++ {
		// TODO change this
		result = processor.ProcessTemplate(processor.HostNameProcessor{}, template)
		result = processor.ProcessTemplate(processor.UserIdProcessor{}, result)
		result = processor.ProcessTemplate(processor.HttpMethodProcessor{}, result)
		result = processor.ProcessTemplate(processor.UrlProcessor{}, result)
		result = processor.ProcessTemplate(processor.StatusCodeProcessor{}, result)
		result = processor.ProcessTemplate(processor.ResourceSizeProcessor{}, result)
		result = processor.ProcessTemplate(processor.RefererProcessor{}, result)
		result = processor.ProcessTemplate(processor.UserAgentProcessor{}, result)
		result = processor.ProcessTemplate(processor.TimestampProcessor{}, result)
		fmt.Println(result)
	}
}
