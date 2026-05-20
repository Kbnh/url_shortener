package sqlite

import (
	"context"
	"fmt"

	"github.com/Kbnh/url_shortener/internal/domain"
)

func (s *Storage) DeleteURL(ctx context.Context, id int64) error {
	query := `DELETE FROM url WHERE id = ?`

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("s.db.ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("res.RowsAffected: %w", err)
	}
	if rows == 0 {
		return domain.ErrURLNotFound
	}

	return nil
}
