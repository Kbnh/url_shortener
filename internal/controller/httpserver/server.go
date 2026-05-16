package httpserver

import (
	"net/http"
	"time"
)

type Config struct {
	Address     string        `env:"ADDR" env-default:"localhost:8080"`
	Timeout     time.Duration `env:"TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
}

type Server struct {
	*http.Server
}

func New(c Config, handler http.Handler) *Server {
	return &Server{
		Server: &http.Server{
			Addr:         c.Address,
			Handler:      handler,
			ReadTimeout:  c.Timeout,
			WriteTimeout: c.Timeout,
			IdleTimeout:  c.IdleTimeout,
		},
	}
}
