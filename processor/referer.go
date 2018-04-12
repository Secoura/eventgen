package processor

import (
	"math/rand"
	"strings"
	"time"

	"github.com/icrowley/fake"
)

type refererProcessor struct{}

var refererKey = "%Referer%"

func (p refererProcessor) process(template string) string {
	if strings.Contains(template, refererKey) {
		referer := getRandomReferer()
		return strings.Replace(template, refererKey, referer, -1)
	}
	return template
}

func getRandomReferer() string {
	referers := []string{"-", "-", "http://" + fake.DomainName(), "-", "-", "-"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return referers[r.Intn(len(referers))]
}