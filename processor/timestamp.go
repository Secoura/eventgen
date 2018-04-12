package processor

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type timestampProcessor struct{}

var timestampKey = "%Timestamp:"

var timestampRegex = "%Timestamp:(.+?)%"

func (p timestampProcessor) process(template string) string {
	regex, _ := regexp.Compile(timestampRegex)
	matches := regex.FindAllStringSubmatch(template, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			startTime := time.Now() // TODO change to use from config
			duration := getRandomTimeInterval(startTime)

			currentTime := startTime.Add(time.Duration(duration)).Format(match[1])
			template = strings.Replace(template, match[0], currentTime, -1)
		}
	}
	return template
}

func getRandomTimeInterval(startTime time.Time) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randMin := r.Intn(15) // TODO change to use from config
	duration := randMin * int(time.Minute)

	// get a random second
	duration = duration - startTime.Second()
	randSec := r.Intn(60)
	duration = duration + (randSec * int(time.Second))

	return duration
}
