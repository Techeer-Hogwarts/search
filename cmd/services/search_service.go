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
func (s *SearchService) PerformSearch(index, query string, limit, offset int) ([]models.CombinedSearchResult, error) {
	return s.repo.Search(index, query, limit, offset)
}

func (s *SearchService) PerformBasicSearch(index, query string, limit, offset int) ([]models.CombinedSearchResult, error) {
	return s.repo.TitleSearch(index, query, limit, offset)
}

func (s *SearchService) PerformUserSearch(index, query string, limit, offset int) ([]models.UserSearchResult, error) {
	return s.repo.UserSearch(query, limit, offset)
}

func (s *SearchService) PerformFinalSearch(index, query string, limit, offset int) ([]models.SearchResult, error) {
	return s.repo.FinalSearch(index, query, limit, offset)
}
