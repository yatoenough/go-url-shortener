package server

import (
	"net/http"

	"github.com/yatoenough/go-url-shortener/internal/config"
	"github.com/yatoenough/go-url-shortener/internal/server/routes"
)

func New(cfg *config.Config) *http.Server {
	return &http.Server{
		Addr:         cfg.Address,
		Handler:      routes.RegisterRoutes(),
		ReadTimeout:  cfg.HTTPServer.RequestTimeout,
		WriteTimeout: cfg.HTTPServer.RequestTimeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
}
