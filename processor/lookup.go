package processor

import (
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type lookupProcessor struct{}

var lookupKey = "%Lookup:"

var lookupRegex = "%Lookup:(.+?)%"

func (p lookupProcessor) process(template string) string {
	regex, _ := regexp.Compile(lookupRegex)
	matches := regex.FindAllStringSubmatch(template, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			value := lookupRandomValue(match[1])
			template = strings.Replace(template, match[0], value, -1)
		}
	}
	return template
}

func lookupRandomValue(file string) string {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}

	data := strings.Split(string(f), "\n")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randLine := r.Intn(len(data))

	return data[randLine]
}
