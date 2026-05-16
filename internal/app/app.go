package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kbnh/url_shortener/config"
	"github.com/Kbnh/url_shortener/internal/adapter/sqlite"
	"github.com/Kbnh/url_shortener/internal/controller/httpserver"
	"github.com/Kbnh/url_shortener/internal/controller/router"
	"github.com/Kbnh/url_shortener/internal/usecase"
)

func Run(ctx context.Context, log *slog.Logger, c config.Config) error {
	storage, err := sqlite.New(c.Sqlite)
	if err != nil {
		log.Error("sqlite.New", slog.String("error", err.Error()))
		return err
	}

	uc := usecase.New(storage)

	router := router.New(log, uc)

	srv := httpserver.New(httpserver.Config(c.HTTPServer), router)

	log.Info("starting server", slog.String("address", c.HTTPServer.Address))

	srvErr := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			srvErr <- fmt.Errorf("server failed: %w", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sig:
		log.Info("shutdown signal recieved")
	case err := <-srvErr:
		return err
	}

	log.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown failed", slog.Any("error", err))
	}

	log.Info("server stopped")

	return nil

}
