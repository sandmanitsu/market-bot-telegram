package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Application Application `yaml:"application" env-required:"true"`
	VK          VK          `yaml:"vkapi" env-required:"true"`
}

type Application struct {
	Host string `yaml:"host" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
	Env  string `yaml:"env" env-required:"true"` // env has two values: "dev" or "prod"
}

type VK struct {
	ClientID    string `yaml:"client_id" env-required:"true"`
	AccessToken string `yaml"access_token" env-required:"true"`
}

var (
	config *Config
	once   sync.Once
)

func MustLoad() *Config {
	if config == nil {
		once.Do(
			func() {
				configPath := filepath.Join("config.yaml")

				if _, err := os.Stat(configPath); err != nil {
					log.Fatalf("Error opening config file: %s", err)
				}

				var newConfig Config
				err := cleanenv.ReadConfig(configPath, &newConfig)
				if err != nil {
					log.Fatalf("Error reading config file: %s", err)
				}

				config = &newConfig
			})
	}

	return config
}
