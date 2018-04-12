package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Duration   time.Duration
	NoOfEvents int
	StartTime  string
}

var config Config

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading config file: " + err.Error())
	}

	config = Config{
		Duration:   parseDuration(v.GetString("duration")),
		NoOfEvents: v.GetInt("noOfEvents"),
		StartTime:  v.GetString("startTime"),
	}
}

func GetConfig() Config {
	return config
}
