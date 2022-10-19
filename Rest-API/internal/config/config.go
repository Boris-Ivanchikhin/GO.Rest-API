package config

import (
	"RestAPI/pkg/logging"
	"sync"

	// *** ilyakaznacheev / cleanenv
	// This is a simple configuration reading tool. It just does the following:
	// -> reads and parses configuration structure from the file
	// -> reads and overwrites configuration structure from environment variables
	// -> writes a detailed variable list to help output

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once

// Config реализован как Singleton
func GetConfig() *Config {

	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
