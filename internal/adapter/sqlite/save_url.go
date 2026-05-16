package sqlite

import (
	"context"
	"fmt"

	"github.com/Kbnh/url_shortener/internal/domain"
	"github.com/mattn/go-sqlite3"
)

func (s *Storage) SaveURL(ctx context.Context, urlToSave, alias string) (int64, error) {
	query := `INSERT INTO url(url, alias) VALUES(?, ?)`

	res, err := s.db.ExecContext(ctx, query, urlToSave, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("query.Exec: %w", domain.ErrURLExists)
		}
		return 0, fmt.Errorf("query.Exec: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Errorf("res.LastInsertId: %w", err)
		return 0, fmt.Errorf("res.LastInsertId: %w", err)
	}

	return id, nil
}
