package httpserver

import "time"

type Config struct {
	Address     string        `env:"ADDR" env-default:"localhost:8080"`
	Timeout     time.Duration `env:"TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
}
