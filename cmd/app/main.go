package main

import (
	"os"

	"github.com/yatoenough/go-url-shortener/internal/config"
	"github.com/yatoenough/go-url-shortener/internal/db/postgres"
	"github.com/yatoenough/go-url-shortener/internal/lib/logger"
	"github.com/yatoenough/go-url-shortener/internal/lib/logger/attrs"
	"github.com/yatoenough/go-url-shortener/internal/server"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)
	log.Info("Starting application...")

	db, err := postgres.New(cfg.ConnectionString)
	if err != nil {
		log.Error("Failed to init db connection", attrs.ErrAttr(err))
		os.Exit(1)
	}
	defer db.Close()

	srv := server.New(cfg, db, log)

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", attrs.ErrAttr(err))
	}
}
