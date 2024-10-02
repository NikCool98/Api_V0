package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

func MustLoad() *Config {
	a := godotenv.Load()
	_ = a
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
