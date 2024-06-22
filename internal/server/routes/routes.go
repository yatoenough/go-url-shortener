package routes

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/yatoenough/go-url-shortener/internal/db/postgres"
	"github.com/yatoenough/go-url-shortener/internal/lib/api/response"
	"github.com/yatoenough/go-url-shortener/internal/server/handlers/url"
)

func RegisterRoutes(storage *postgres.Storage, logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	urlHandler := url.New(storage, logger)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, response.Error(http.StatusNotFound, "Route not found."))
	})
	r.Post("/shorten", urlHandler.Shorten)
	r.Get("/{alias}", urlHandler.Redirect)
	return r
}
