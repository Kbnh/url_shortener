package usecase

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/Kbnh/url_shortener/config"
	"github.com/Kbnh/url_shortener/internal/adapter/http"
	"github.com/Kbnh/url_shortener/internal/adapter/storage"
	"github.com/Kbnh/url_shortener/internal/connector"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(ctx context.Context, log *slog.Logger, c config.Config) {
	storage, err := storage.New(c.Sqlite)
	if err != nil {
		log.Error("sqllite.New", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		connector.New(log),
		middleware.Recoverer,
	)

	router.Post("/url", http.SaveURL(log, storage))

	log.Info("starting server", slog.String("address", c.HTTPServer.Address))

	srv := &http.Server{
		Addr:         c.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  c.HTTPServer.Timeout,
		WriteTimeout: c.HTTPServer.Timeout,
		IdleTimeout:  c.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")

}
