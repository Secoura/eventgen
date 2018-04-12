package processor

import (
	"strings"
)

type Processor interface {
	process(template string) string
}

var processors map[string]Processor

func RegisterProcessors() {
	processors = make(map[string]Processor)
	processors[hostNameKey] = hostNameProcessor{}
	processors[httpMethodKey] = httpMethodProcessor{}
	processors[refererKey] = refererProcessor{}
	processors[resourceSizeKey] = resourceSizeProcessor{}
	processors[statusCodeKey] = statusCodeProcessor{}
	processors[timestampKey] = timestampProcessor{}
	processors[urlKey] = urlProcessor{}
	processors[userAgentKey] = userAgentProcessor{}
	processors[userIdKey] = userIdProcessor{}
}

func ProcessTemplate(template string) string {
	result := template
	for key, p := range processors {
		if strings.Contains(result, key) {
			result = p.process(result)
		}
	}
	return result
}
