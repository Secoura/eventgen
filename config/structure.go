package config

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type Config struct {
	Template   string        `json:"template"`
	Duration   time.Duration `json:"duration"`
	NoOfEvents int           `json:"noOfEvents"`
	StartTime  string        `json:"startTime"`
}

func (c *Config) UnmarshalJSON(b []byte) error {
	var conf map[string]string
	if err := json.Unmarshal(b, &conf); err != nil {
		return err
	}
	if !checkRequiredConfigKeys(conf) {
		return errors.New("missing required config key")
	}

	c.Template = conf["template"]
	c.StartTime = conf["startTime"]
	c.Duration = parseDuration(conf["duration"])

	noOfEvents, err := strconv.Atoi(conf["noOfEvents"])
	if err != nil {
		return err
	}
	c.NoOfEvents = noOfEvents
	return nil
}

func checkRequiredConfigKeys(conf map[string]string) bool {
	if _, exists := conf["noOfEvents"]; !exists {
		return false
	}

	_, exists := conf["template"]
	return exists
}
