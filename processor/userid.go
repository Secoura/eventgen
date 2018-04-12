package processor

import (
	"strings"
)

type userIdProcessor struct{}

var userIdKey = "%UserId%"

func (p userIdProcessor) process(template string) string {
	if strings.Contains(template, userIdKey) {
		return strings.Replace(template, userIdKey, "-", -1)
	}
	return template
}
