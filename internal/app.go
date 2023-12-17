package grpcapp

import (
	"log/slog"
	"main/internal/app"
	auth "main/internal/domain/service"
	"main/internal/storage"
	"time"
)

type App struct {
	Grpc *app.App
}

func New(logger slog.Logger, port string, host string, user string, password string, expToken time.Duration) *App {
	storage, err := storage.New(logger, host, user, password, port)
	if err != nil {
		logger.Error("Could not start new app")
	}
	authService := auth.New(&logger, expToken, storage, storage)

	grpcApp := app.New(logger, port, expToken, authService)

	return &App{Grpc: grpcApp}
}
