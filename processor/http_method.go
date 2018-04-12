package processor

import (
	"math/rand"
	"strings"
	"time"
)

type httpMethodProcessor struct{}

var httpMethodKey = "%HttpMethod%"

func (p httpMethodProcessor) process(template string) string {
	if strings.Contains(template, httpMethodKey) {
		method := getRandomHttpMethod()
		return strings.Replace(template, httpMethodKey, method, -1)
	}
	return template
}

func getRandomHttpMethod() string {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return methods[r.Intn(len(methods))]
}
