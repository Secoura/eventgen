package processor

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/yeoji/eventgen/config"
)

type timestampProcessor struct{}

var timestampKey = "%Timestamp:"

var timestampRegex = "%Timestamp:(.+?)%"

var startTimeFormat = "02/Jan/2006:15:04:05 -0700"

func (p timestampProcessor) process(template string) string {
	regex, _ := regexp.Compile(timestampRegex)
	matches := regex.FindAllStringSubmatch(template, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			startTime, err := time.Parse(startTimeFormat, config.GetConfig().StartTime)
			if err != nil {
				startTime = time.Now()
			}
			duration := getRandomTimeInterval(startTime)

			currentTime := startTime.Add(duration).Format(match[1])
			template = strings.Replace(template, match[0], currentTime, -1)
		}
	}
	return template
}

func getRandomTimeInterval(startTime time.Time) time.Duration {
	configDuration := config.GetConfig().Duration

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randMin := r.Intn(int(configDuration.Minutes()))
	duration := randMin * int(time.Minute)

	// get a random second
	duration = duration - startTime.Second()
	randSec := r.Intn(60)
	duration = duration + (randSec * int(time.Second))

	return time.Duration(duration)
}
