package usecase

import (
	"context"
	"fmt"
)

func (u *UseCase) DeleteURL(ctx context.Context, id int64) error {
	err := u.repo.DeleteURL(ctx, id)
	if err != nil {
		return fmt.Errorf("DeleteURL: %w", err)
	}
	return nil
}
