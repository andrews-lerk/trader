package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

const DefaultConfigPath = "./config/config.yaml"

type Config struct {
	Api    Api    `yaml:"api"`
	Logger Logger `yaml:"logger"`
}

type Api struct {
	Host    string        `yaml:"host" env-default:"localhost"`
	Port    int           `yaml:"port" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type Logger struct {
	Json  bool   `yaml:"json" env-default:"false"`
	Level string `yaml:"level" env-default:"DEBUG"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = DefaultConfigPath
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config file: %s", err)
	}
	return &cfg
}
