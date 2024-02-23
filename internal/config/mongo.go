package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type MongoConfig struct {
	DbCommonConnectString string `env:"DB_COMMON_CONNECT_STRING" required:"true"`
}

func GetMongoConfig() *MongoConfig {

	// err := godotenv.Load()

	// Debug code
	path, _ := os.Executable()
	path = filepath.Join(path, "../../.env")
	err := godotenv.Load(path)
	//

	if err != nil {
		log.Fatalf("Mongo unable to load .env file: %e", err)
	}

	cfg := MongoConfig{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	return &cfg
}
