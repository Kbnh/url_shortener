package main

import (
	"context"
	"log/slog"

	"github.com/Kbnh/url_shortener/config"
	"github.com/Kbnh/url_shortener/internal/usecase"
	"github.com/Kbnh/url_shortener/pkg/logger"
)

func main() {
	cfg := config.MustLoad() // init config

	log := logger.Init(cfg.Logger) // init logger
	log = log.With(slog.String("app", cfg.App.Name), slog.String("version", cfg.App.Version))

	ctx := context.Background()

	usecase.Run(ctx, log, *cfg)
	// TODO: app run:
	// 			init storage
	// 			init router
	// 			run server

}
