package usecase

import "context"

type Repo interface {
	SaveURL(ctx context.Context, urlToSave, alias string) (int64, error)
	GetURL(ctx context.Context, alias string) (string, error)
	DeleteURL(ctx context.Context, id int64) error
}

type UseCase struct {
	repo Repo
}

func New(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

const aliasLength = 6
