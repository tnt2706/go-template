package config

import (
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type GrpcConfig struct {
	CalculatorServer string `env:"CALCULATOR_SERVER" required:"true"`
}

func GetGrpcConfig() *GrpcConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("gRPC unable to load .env file: %e", err)
	}

	cfg := GrpcConfig{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	return &cfg
}
