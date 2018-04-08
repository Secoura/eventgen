package processor

import (
	"strings"

	"github.com/icrowley/fake"
)

type UserAgentProcessor struct{}

var userAgentKey = "%UserAgent%"

func (p UserAgentProcessor) process(template string) string {
	if strings.Contains(template, userAgentKey) {
		return strings.Replace(template, userAgentKey, fake.UserAgent(), -1)
	}
	return template
}
