package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Path string `env:"PATH" env-required:"true"`
}

type Storage struct {
	db *sql.DB
}

func New(c Config) (*Storage, error) {
	db, err := sql.Open("sqlite3", c.Path)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	db.SetMaxOpenConns(1)

	return &Storage{db: db}, nil
}
