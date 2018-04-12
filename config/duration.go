package config

import (
	"regexp"
	"strconv"
	"time"
)

var durationRegex = "(?:(\\d+)d\\s?)?(?:(\\d+)h\\s?)?(?:(\\d+)m\\s?)?"

func parseDuration(duration string) time.Duration {
	regex, _ := regexp.Compile(durationRegex)
	matches := regex.FindAllStringSubmatch(duration, -1)
	if len(matches) > 0 {
		match := matches[0]
		day, _ := strconv.Atoi(match[1])
		hour, _ := strconv.Atoi(match[2])
		min, _ := strconv.Atoi(match[3])

		return createDuration(day, hour, min)
	}
	return time.Duration(15 * time.Minute)
}

func createDuration(day int, hour int, min int) time.Duration {
	duration := day * int(time.Hour) * 24
	duration += hour * int(time.Hour)
	duration += min * int(time.Minute)
	return time.Duration(duration)
}
