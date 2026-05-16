package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	Env string `env:"ENV" env-required:"true"`
}

var (
	envLocal = "local"
	envDev   = "dev"
)

func Init(c Config) *slog.Logger {
	var log *slog.Logger

	switch c.Env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}
