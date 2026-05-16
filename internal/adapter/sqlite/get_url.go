package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Kbnh/url_shortener/internal/domain"
)

func (s *Storage) GetURL(ctx context.Context, alias string) (string, error) {
	query := `SELECT url FROM url WHERE alias = ?`

	var res string

	if err := s.db.QueryRowContext(ctx, query, alias).Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("query.QueryRow: %w", domain.ErrURLNotFound)
		}
		return "", fmt.Errorf("query.QueryRow: %w", err)
	}

	return res, nil
}
