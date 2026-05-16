package usecase

import (
	"context"
	"fmt"
)

func (u *UseCase) GetURL(ctx context.Context, alias string) (string, error) {
	res, err := u.repo.GetURL(ctx, alias)
	if err != nil {
		return "", fmt.Errorf("u.repo.GetURL: %w", err)
	}

	return res, nil
}
