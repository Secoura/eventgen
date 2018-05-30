package config

import (
	"encoding/json"
	"log"

	"github.com/spf13/viper"
	"os"
)

var config Config

func LoadConfig() {
	if success := loadPluginConfig(); !success {
		loadFileConfig()
	}
}

func loadPluginConfig() bool {
	pluginConf := os.Getenv("PLUGIN_CONFIG")
	if pluginConf == "" {
		return false
	}

	config = Config{}
	if err := json.Unmarshal([]byte(pluginConf), &config); err != nil {
		log.Printf("Failed to parse PLUGIN_CONFIG from environment variable: %v", err)
		return false
	}
	return true
}

func loadFileConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading config file: " + err.Error())
	}
	config = Config{
		Template:   viper.GetString("template"),
		Duration:   parseDuration(viper.GetString("duration")),
		NoOfEvents: viper.GetInt("noOfEvents"),
		StartTime:  viper.GetString("startTime"),
	}
}

func GetConfig() Config {
	return config
}
