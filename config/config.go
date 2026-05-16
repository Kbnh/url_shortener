package config

import (
	"fmt"
	"log"

	"github.com/Kbnh/url_shortener/internal/adapter/sqlite"
	"github.com/Kbnh/url_shortener/internal/controller/httpserver"
	"github.com/Kbnh/url_shortener/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type App struct { // Конфигурация приложения
	Name    string `env:"NAME"    env-required:"true"`
	Version string `env:"VERSION" env-required:"true"`
}

type Config struct {
	App        App               `env-prefix:"APP_"`
	Sqlite     sqlite.Config     `env-prefix:"DB_"`
	HTTPServer httpserver.Config `env-prefix:"HTTP_"`
	Logger     logger.Config     `env-prefix:"LOGGER_"`
}

func New() (Config, error) {
	var c Config

	_ = godotenv.Load()

	if err := cleanenv.ReadEnv(&c); err != nil {
		return Config{}, fmt.Errorf("cleanenv.ReadEnv: %w", err)
	}

	return c, nil
}

func MustLoad() *Config {
	c, err := New()
	if err != nil {
		log.Fatalf("config.MustLoad: %w", err)
	}
	return &c
}
