package processor

import (
	"strings"

	"github.com/icrowley/fake"
)

type UrlProcessor struct{}

var urlKey = "%Url%"

func (p UrlProcessor) process(template string) string {
	if strings.Contains(template, urlKey) {
		// TODO fix this later
		fakeResource := strings.Join(strings.Split(fake.Product(), " "), "-")
		return strings.Replace(template, urlKey, "/" + fakeResource, -1)
	}
	return template
}

