package processor

import (
	"math/rand"
	"strings"
	"time"
)

type statusCodeProcessor struct{}

var statusCodeKey = "%StatusCode%"

func (p statusCodeProcessor) process(template string) string {
	if strings.Contains(template, statusCodeKey) {
		statusCode := getRandomStatusCode()
		return strings.Replace(template, statusCodeKey, statusCode, -1)
	}
	return template
}

func getRandomStatusCode() string {
	codes := []string{"200", "404", "500"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return codes[r.Intn(len(codes))]
}
