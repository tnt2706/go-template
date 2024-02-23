package config

import (
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type ContainerConfig struct {
	Port      string `env:"PORT" default:"8080"`
	HostName  string `env:"HOST_NAME" default:"lineapp"`
	Namespace string `env:"NAMESPACE" default:"alpha"`
	NodeEnv   string `env:"NODE_ENV" default:"development"`
}

func GetContainerConfig() *ContainerConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Container unable to load .env file: %e", err)
	}

	cfg := ContainerConfig{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	return &cfg
}
