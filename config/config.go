package config

import (
	"encoding/json"
	"log"

	"github.com/spf13/viper"
)

var config Config

func LoadConfig() {
	err := loadPluginConfig()
	if err != nil {
		log.Println("Could not read config from env variable: " + err.Error())
		log.Println("Reading config from file...")
		loadFileConfig()
	}
}

func loadPluginConfig() error {
	viper.SetEnvPrefix("plugin")
	viper.BindEnv("config")

	pluginConf := viper.GetString("config")
	config = Config{}
	return json.Unmarshal([]byte(pluginConf), &config)
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
