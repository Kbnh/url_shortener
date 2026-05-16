package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/url_shortener/internal/seeder"
)

func (u *UseCase) SaveURL(ctx context.Context, url, alias string) (int64, error) {
	if alias == "" {
		alias = seeder.NewRandomString(aliasLength)
	}

	id, err := u.repo.SaveURL(ctx, url, alias)
	if err != nil {
		return 0, fmt.Errorf("u.repo.SaveURL: %w", err)
	}

	return id, nil
}
