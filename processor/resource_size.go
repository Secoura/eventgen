package processor

import (
	"strings"

	"github.com/icrowley/fake"
)

type ResourceSizeProcessor struct{}

var resourceSizeKey = "%ResourceSize%"

func (p ResourceSizeProcessor) process(template string) string {
	if strings.Contains(template, resourceSizeKey) {
		return strings.Replace(template, resourceSizeKey, fake.Digits(), -1)
	}
	return template
}
