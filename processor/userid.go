package processor

import (
	"strings"
)

type UserIdProcessor struct{}

var userIdKey = "%UserId%"

func (p UserIdProcessor) process(template string) string {
	if strings.Contains(template, userIdKey) {
		return strings.Replace(template, userIdKey, "-", -1)
	}
	return template
}
