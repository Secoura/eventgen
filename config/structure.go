package config

import (
	"encoding/json"
	"errors"
	"time"
)

type Config struct {
	Template       string        `json:"template"`
	Delimiter      string        `json:"delimiter"`
	Duration       time.Duration `json:"duration"`
	NumberOfEvents int           `json:"number_of_events"`
	StartTime      string        `json:"start_time"`
}

func (c *Config) UnmarshalJSON(b []byte) error {
	var conf map[string]interface{}
	if err := json.Unmarshal(b, &conf); err != nil {
		return err
	}
	if !checkRequiredConfigKeys(conf) {
		return errors.New("missing required config key")
	}

	c.Template = conf["template"].(string)
	if startTime, ok := conf["start_time"]; ok {
		c.StartTime = startTime.(string)
	}
	if delimiter, ok := conf["delimiter"]; ok {
		c.Delimiter = delimiter.(string)
	}

	switch t := conf["duration"].(type) {
	case string:
		c.Duration = parseDuration(t)
	case time.Duration:
		c.Duration = t
	default:
		return errors.New("invalid 'duration' value")
	}

	noOfEvents, ok := conf["number_of_events"].(float64)
	if !ok {
		return errors.New("invalid 'number_of_events' value")
	}
	c.NumberOfEvents = int(noOfEvents)
	return nil
}

func checkRequiredConfigKeys(conf map[string]interface{}) bool {
	if _, exists := conf["number_of_events"]; !exists {
		return false
	}

	_, exists := conf["template"]
	return exists
}
