package repositories

import "github.com/Techeer-Hogwarts/search/cmd/models"

type Repository interface {
	CombinedSearch(query string, limit, offset int) ([]models.SearchResult, error)
	Search(index, query string, limit, offset int) ([]models.SearchResult, error)
}
