package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	Timeout       string
	TokenLifetime time.Duration
}

func New(logger slog.Logger) *Config {
	if err := godotenv.Load("./internal/config/config.env"); err != nil {
		logger.Error("Could not find or parse .env file")
	}
	port := checkFilledEnv("GRPC_PORT", logger)
	timeout := checkFilledEnv("TIMEOUT", logger)
	lifetime := checkFilledEnv("TOKEN_LIFETIME", logger)
	lifeTimeDuration, err := time.ParseDuration(lifetime)
	if err != nil {
		fmt.Println("Error while parsing time duration")
	}

	cfg := Config{Port: port, Timeout: timeout, TokenLifetime: lifeTimeDuration}
	return &cfg
}
func checkFilledEnv(envParam string, logger slog.Logger) string {
	envVar, exist := os.LookupEnv(envParam)
	if !exist {
		logger.Error("Could not find", slog.String("param", envParam))
	}
	return envVar
}
