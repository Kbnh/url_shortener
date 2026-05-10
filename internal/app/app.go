package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/Kbnh/url_shortener/config"
	v1 "github.com/Kbnh/url_shortener/internal/controller/v1"
	mwLogger "github.com/Kbnh/url_shortener/internal/http_server/middleware/logger"
	"github.com/Kbnh/url_shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(ctx context.Context, log *slog.Logger, c config.Config) {
	storage, err := sqlite.New(c.Sqlite)
	if err != nil {
		log.Error("sqllite.New", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		mwLogger.New(log),
		middleware.Recoverer,
	)

	router.Post("/url", v1.SaveURL(log, storage))

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
