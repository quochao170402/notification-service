package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type MongoConfig struct {
	MongoURI  string
	MongoDB   string
	MongoUser string
	MongoPass string
}

type AppConfig struct {
	Port string
}

type Config struct {
	MongoConfig MongoConfig
	AppConfig   AppConfig
}

// Load reads environment variables from .env and system env.
func Load() *Config {
	// Load .env if it exists (does nothing if missing)
	_ = godotenv.Load()

	mongoCfg := MongoConfig{
		MongoURI:  os.Getenv("MONGO_URI"),
		MongoDB:   os.Getenv("MONGO_DB"),
		MongoUser: os.Getenv("MONGO_USER"),
		MongoPass: os.Getenv("MONGO_PASS"),
	}

	if mongoCfg.MongoURI == "" || mongoCfg.MongoDB == "" {
		log.Fatal("Missing MONGO_URI or MONGO_DB in environment")
	}

	appConfig := AppConfig{
		Port: os.Getenv("APP_PORT"),
	}

	if appConfig.Port == "" {
		appConfig.Port = "8080" // fallback
	}

	return &Config{
		MongoConfig: mongoCfg,
		AppConfig:   appConfig,
	}
}
