package search

import "context"

type (
	Search interface {
		Search(word string, ctx context.Context) ([]string, error)
	}

	Repository interface {
		GetSearch(word string, ctx context.Context) ([]string, error)
	}

	API interface {
		Search(word string) ([]string, error)
	}
)
