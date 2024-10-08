package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	DB         ConfigDB   `yaml:"db"`
	HTTPServer ConfigHttp `yaml:"http_server"`
}

type ConfigHttp struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type ConfigDB struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

func MustLoad(configPath string) *Config {

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
