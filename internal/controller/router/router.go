package router

import (
	"context"
	"log/slog"

	"github.com/Kbnh/url_shortener/internal/controller/httpserver"
	"github.com/Kbnh/url_shortener/internal/controller/httpserver/middleware/logger"
	v1 "github.com/Kbnh/url_shortener/internal/controller/httpserver/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type UseCase interface {
	SaveURL(ctx context.Context, url, alias string) (int64, error)
	GetURL(ctx context.Context, alias string) (string, error)
	DeleteURL(ctx context.Context, id int64) error
}

func New(log *slog.Logger, uc UseCase, cfg httpserver.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		logger.New(log),
		middleware.Recoverer,
	)

	r.Route("/url", func(router chi.Router) {
		router.Use(middleware.BasicAuth("url-shortener", map[string]string{
			cfg.User: cfg.Password,
		}))
		router.Post("/", v1.SaveURL(log, uc))
		router.Delete("/{id}", v1.DeleteURL(log, uc))
	})

	r.Get("/{alias}", v1.GetURL(log, uc))

	return r
}
