package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Kbnh/url_shortener/internal/domain"
)

func (s *Storage) GetURL(alias string) (string, error) {
	query, err := s.db.Prepare(`
	SELECT url FROM url
	WHERE alias = ?
	`)
	if err != nil {
		return "", fmt.Errorf("s.db.Prepare: %w", err)
	}

	var res string
	if err = query.QueryRow(alias).Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("query.QueryRow: %w", domain.ErrURLNotFound)
		}
		return "", fmt.Errorf("query.QueryRow: %w", err)
	}

	return res, nil
}
