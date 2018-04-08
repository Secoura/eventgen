package processor

import (
	"math/rand"
	"strings"
	"time"
)

type StatusCodeProcessor struct{}

var statusCodeKey = "%StatusCode%"

func (p StatusCodeProcessor) process(template string) string {
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
