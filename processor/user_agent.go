package processor

import (
	"strings"

	"github.com/icrowley/fake"
)

type userAgentProcessor struct{}

var userAgentKey = "%UserAgent%"

func (p userAgentProcessor) process(template string) string {
	if strings.Contains(template, userAgentKey) {
		return strings.Replace(template, userAgentKey, fake.UserAgent(), -1)
	}
	return template
}
