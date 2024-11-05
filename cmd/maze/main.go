package main

import (
	"log/slog"
	"os"

	"github.com/KazikovAP/maze/config"
	"github.com/KazikovAP/maze/internal/application"
	"github.com/KazikovAP/maze/internal/infrastructure"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := config.NewConfig()
	cfg.Init()

	ioAdapter := infrastructure.NewIOAdapter(os.Stdin, os.Stdout, logger)

	app := application.NewApp(cfg, ioAdapter)
	if err := app.Start(); err != nil {
		logger.Error("Application failed to start", "error", err)
	}
}
