package rpcLogger

import (
	"fmt"
	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	if logger == nil {
		fmt.Println("Could not create logger")
	}
	return logger
}
