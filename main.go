package main

import (
	rpcLogger "main/common"
	"main/internal/app"
	"main/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := rpcLogger.SetupLogger()
	cfg := config.New(*logger)
	application := app.New(*logger, cfg.Port, cfg.TokenLifetime)

	application.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop()
	//TODO собрать весь jrpc тут
}
