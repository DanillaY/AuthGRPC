package main

import (
	rpcLogger "main/common"
	grpcapp "main/internal"

	"main/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := rpcLogger.SetupLogger()
	cfg := config.New(*logger)
	application := grpcapp.New(*logger, cfg.Port, cfg.Host, cfg.User, cfg.Password, cfg.TokenLifetime)
	application.Grpc.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Grpc.Stop()
}
