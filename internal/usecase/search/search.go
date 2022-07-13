package search

import "context"

type UseCase struct {
	repo Repository
}

func NewUseCase(repository Repository) *UseCase {
	return &UseCase{
		repo: repository,
	}
}

func (uc *UseCase) Search(ctx context.Context, word string) ([]string, error) {
	if result, err := uc.repo.GetSearch(word, ctx); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
