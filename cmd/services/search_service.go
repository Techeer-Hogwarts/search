package services

import (
	"github.com/Techeer-Hogwarts/search/cmd/models"
	"github.com/Techeer-Hogwarts/search/cmd/repositories"
)

type SearchService struct {
	repo repositories.Repository
}

// NewSearchService initializes the service
func NewSearchService(repo repositories.Repository) *SearchService {
	return &SearchService{repo: repo}
}

// PerformSearch executes the search query
func (s *SearchService) PerformSearch(query string, limit, offset int) ([]models.SearchResult, error) {
	return s.repo.Search(query, limit, offset)
}
