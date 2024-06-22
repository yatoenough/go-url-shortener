package main

import (
	"log/slog"
	"os"

	"github.com/yatoenough/go-url-shortener/internal/config"
	"github.com/yatoenough/go-url-shortener/internal/db/postgres"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "production"
)

func main() {
	cfg := config.MustLoad()

	log := initLogger(cfg.Env)
	log.Info("Starting application...")

	db, err := postgres.New(cfg.ConnectionString)
	if err != nil {
		log.Error("Failed to init db connection", errAttr(err))
		os.Exit(1)
	}

	_ = db
}

func initLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}

func errAttr(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
