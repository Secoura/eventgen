package processor

import (
	"strings"

	"github.com/icrowley/fake"
)

type resourceSizeProcessor struct{}

var resourceSizeKey = "%ResourceSize%"

func (p resourceSizeProcessor) process(template string) string {
	if strings.Contains(template, resourceSizeKey) {
		return strings.Replace(template, resourceSizeKey, fake.Digits(), -1)
	}
	return template
}
