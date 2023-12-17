package app

import (
	"fmt"
	"log/slog"
	"main/internal/server"
	"net"
	"time"

	"google.golang.org/grpc"
)

type App struct {
	logger *slog.Logger
	grpc   *grpc.Server
	port   string
}

func New(logger slog.Logger, port string, expToken time.Duration, auth server.IAuth) *App {
	grpc := grpc.NewServer()

	server.RegisterServer(grpc, auth)
	return &App{logger: &logger, grpc: grpc, port: port}
}
func (a *App) Run() error {
	log := a.logger.With(slog.String("port", a.port))
	log.Info("Starting grpc server")

	l, err := net.Listen("tcp", a.port)

	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := a.grpc.Serve(l); err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Info("Server is running ", slog.String("addres", l.Addr().String()))
	return nil
}

func (a *App) Stop() {
	a.logger.Info("Stopping server")
	a.grpc.GracefulStop()
}
