package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Kbnh/url_shortener/config"
	"github.com/Kbnh/url_shortener/internal/app"
	"github.com/Kbnh/url_shortener/pkg/logger"
)

func main() {
	cfg := config.MustLoad() // init config

	log := logger.Init(cfg.Logger) // init logger
	log = log.With(slog.String("app", cfg.App.Name), slog.String("version", cfg.App.Version))

	ctx := context.Background()

	err := app.Run(ctx, log, *cfg)
	if err != nil {
		log.Error("failed app run")
		os.Exit(1)
	}

}
