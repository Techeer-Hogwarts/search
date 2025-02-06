package repositories

import "github.com/Techeer-Hogwarts/search/cmd/models"

type Repository interface {
	Search(query string, limit, offset int) ([]models.SearchResult, error)
}
