package router

import (
	"context"
	"log/slog"

	"github.com/Kbnh/url_shortener/internal/controller/httpserver/middleware/logger"
	v1 "github.com/Kbnh/url_shortener/internal/controller/httpserver/v1"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type UseCase interface {
	SaveURL(ctx context.Context, url, alias string) (int64, error)
	GetURL(ctx context.Context, alias string) (string, error)
}

func New(log *slog.Logger, uc UseCase) *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		logger.New(log),
		middleware.Recoverer,
	)

	r.Post("/url", v1.SaveURL(log, uc))
	r.Get("/{alias}", v1.GetURL(log, uc))

	return r
}
