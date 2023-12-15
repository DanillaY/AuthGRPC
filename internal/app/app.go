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

func New(logger slog.Logger, port string, expToken time.Duration) *App {
	grpc := grpc.NewServer()
	//TODO: убрать интерфейс-затычку
	var auth server.IAuth
	server.RegisterServer(grpc, auth)
	return &App{logger: &logger, grpc: grpc, port: port}
}
func (a *App) Run() error {
	log := a.logger
	log.Info("Starting grpc server")

	l, err := net.Listen("tcp", a.port)

	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := a.grpc.Serve(l); err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Info("Server is running ", slog.String("addres", l.Addr().String()))
	log.Info("test", l.Addr().String())

	return nil
}

func (a *App) Stop() {
	a.logger.Info("Stopping server")
	a.grpc.GracefulStop()
}
