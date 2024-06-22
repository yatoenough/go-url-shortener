package server

import (
	"log/slog"
	"net/http"

	"github.com/yatoenough/go-url-shortener/internal/config"
	"github.com/yatoenough/go-url-shortener/internal/db/postgres"
	"github.com/yatoenough/go-url-shortener/internal/server/routes"
)

func New(cfg *config.Config, storage *postgres.Storage, logger *slog.Logger) *http.Server {
	return &http.Server{
		Addr:         cfg.Address,
		Handler:      routes.RegisterRoutes(storage, logger),
		ReadTimeout:  cfg.HTTPServer.RequestTimeout,
		WriteTimeout: cfg.HTTPServer.RequestTimeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
}
