package config

import (
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort         string `env:"SERVER_PORT,required"`
	DBHost             string `env:"DB_HOST,required"`
	DBName             string `env:"DB_NAME,required"`
	DBUser             string `env:"DB_USER,required"`
	DBPassword         string `env:"DB_PASSWORD,required"`
	DBSSLMode          string `env:"DB_SSLMODE,required"`
	AccessTokenSecret  string `env:"ACCESS_TOKEN_SECRET,required"`
	RefreshTokenSecret string `env:"REFRESH_TOKEN_SECRET,required"`
	AccessTokenExpiry  int    `env:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry int    `env:"REFRESH_TOKEN_EXPIRY"`
	RedisHost          string `env:"REDIS_HOST,required"`
	RedisPort          string `env:"REDIS_PORT,required"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Unable to load .env: %e", err)
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from .env: %e", err)
	}

	return config
}
