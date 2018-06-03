package processor

import (
	"strings"

	"math/rand"
	"strconv"
)

type resourceSizeProcessor struct{}

var resourceSizeKey = "%ResourceSize%"

func (p resourceSizeProcessor) process(template string) string {
	if strings.Contains(template, resourceSizeKey) {
		fakeNumber := rand.Intn(10485760) + 1
		return strings.Replace(template, resourceSizeKey, strconv.Itoa(fakeNumber), -1)
	}
	return template
}
