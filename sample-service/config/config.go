package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port    string `envconfig:"port" default:"9001" validate:"numeric"`
	Address string `envconfig:"address" default:""`
}

func NewConfig() *Config {
	var conf Config
	err := envconfig.Process("ACCOUNT", &conf)
	if err != nil {
		log.Fatalf("fail to proceed the config: %v", err)
	}
	return &conf
}
