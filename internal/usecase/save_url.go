package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/url_shortener/internal/seeder"
)

func (u *UseCase) SaveURL(ctx context.Context, url, alias string) (string, error) {
	if alias == "" {
		alias = seeder.NewRandomString(aliasLenght)
	}

	_, err := u.repo.SaveURL(ctx, url, alias)
	if err != nil {
		return "", fmt.Errorf("u.repo.SaveURL: %w", err)
	}

	return alias, nil
}
