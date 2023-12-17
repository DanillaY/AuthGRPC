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
	Host          string
	User          string
	Password      string
	TokenLifetime time.Duration
}

func New(logger slog.Logger) *Config {
	if err := godotenv.Load("./internal/config/config.env"); err != nil {
		logger.Error("Could not find or parse .env file")
	}
	checkEnvFle("./internal/config/config.env", logger)
	checkEnvFle("./config.env", logger)

	port := checkFilledEnv("GRPC_PORT", logger)
	timeout := checkFilledEnv("TIMEOUT", logger)
	lifetime := checkFilledEnv("TOKEN_LIFETIME", logger)
	host := checkFilledEnv("POSTGRES_HOSTNAME", logger)
	user := checkFilledEnv("POSTGRES_USER", logger)
	password := checkFilledEnv("POSTGRES_PASSWORD", logger)

	lifeTimeDuration, err := time.ParseDuration(lifetime)
	if err != nil {
		fmt.Println("Error while parsing time duration")
	}

	cfg := Config{Port: port, Timeout: timeout, TokenLifetime: lifeTimeDuration, Host: host, User: user, Password: password}
	return &cfg
}
func checkEnvFle(path string, logger slog.Logger) {
	if err := godotenv.Load(path); err != nil {
		logger.Error("Could not find or parse .env file")
	}
}
func checkFilledEnv(envParam string, logger slog.Logger) string {
	envVar, exist := os.LookupEnv(envParam)
	if !exist {
		logger.Error("Could not find", slog.String("param", envParam))
	}
	return envVar
}
