package routes

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yatoenough/go-url-shortener/internal/db/postgres"
	"github.com/yatoenough/go-url-shortener/internal/server/handlers/url"
)

func RegisterRoutes(storage *postgres.Storage, logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	urlHandler := url.New(storage, logger)

	r.Post("/api/shorten", urlHandler.Shorten)
	return r
}
