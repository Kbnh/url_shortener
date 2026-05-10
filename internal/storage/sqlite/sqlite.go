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

	query, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL);	
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);	
	`)
	if err != nil {
		return nil, fmt.Errorf("db.Prepare: %w", err)
	}

	_, err = query.Exec()
	if err != nil {
		return nil, fmt.Errorf("query.Exec: %w", err)
	}

	return &Storage{db: db}, nil
}
