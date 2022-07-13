package repo

import (
	"context"
	"reegle/pkg/dict"
)

type Repository struct {
	db *dict.WordDB
}

func NewSearchRepo(db *dict.WordDB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetSearch(word string, ctx context.Context) ([]string, error) {
	res, err := r.db.SearchByRegex(word)
	if err != nil {
		return nil, err
	}
	return *res, nil
}
