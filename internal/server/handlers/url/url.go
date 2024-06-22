package url

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/yatoenough/go-url-shortener/internal/db"
	"github.com/yatoenough/go-url-shortener/internal/db/postgres"
	"github.com/yatoenough/go-url-shortener/internal/lib/api/response"
	"github.com/yatoenough/go-url-shortener/internal/lib/logger/attrs"
	"github.com/yatoenough/go-url-shortener/internal/lib/model/url/dto"
	"github.com/yatoenough/go-url-shortener/internal/lib/util"
)

type UrlsHandler struct {
	storage *postgres.Storage
	logger  *slog.Logger
}

func New(storage *postgres.Storage, logger *slog.Logger) *UrlsHandler {
	return &UrlsHandler{
		storage: storage,
		logger:  logger,
	}
}

func (uh *UrlsHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	var body dto.URLRequest

	err := render.DecodeJSON(r.Body, &body)
	if err != nil {
		uh.logger.Info("failed to decode request", attrs.ErrAttr(err))
		render.JSON(w, r, response.Error(http.StatusBadRequest, "Failed to decode request."))
		return
	}

	if err := validator.New().Struct(body); err != nil {
		uh.logger.Info("invalid url", attrs.ErrAttr(err))
		render.JSON(w, r, response.Error(http.StatusBadRequest, "Invalid URL."))
		return
	}

	alias := util.NewRandomString(6)
	_, err = uh.storage.SaveURL(body.URL, alias)
	if err != nil {
		uh.logger.Error("error while saving to DB", attrs.ErrAttr(err))
		render.JSON(w, r, response.Error(http.StatusInternalServerError, "Internal server error."))
		return
	}

	render.JSON(w, r, dto.URLResponse{Response: response.OK(), Alias: alias})
}

func (uh *UrlsHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")

	resURL, err := uh.storage.GetURL(alias)
	if err != nil {
		if errors.Is(err, db.ErrURLNotFound) {
			uh.logger.Info("url not found")
			render.JSON(w, r, response.Error(http.StatusBadRequest, "URL not found."))
			return
		}

		uh.logger.Error("failed to get url")
		render.JSON(w, r, response.Error(http.StatusInternalServerError, "Internal server error."))
		return
	}

	http.Redirect(w, r, resURL, http.StatusPermanentRedirect)
}
