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

func (s *SearchService) PerformUserSearch(query string, limit, offset int) ([]models.UserSearchResult, error) {
	return s.repo.UserSearch(query, limit, offset)
}

func (s *SearchService) PerformProjectSearch(query string, limit, offset int) ([]models.ProjectSearchResult, error) {
	return s.repo.ProjectSearch(query, limit, offset)
}

func (s *SearchService) PerformStudySearch(query string, limit, offset int) ([]models.StudySearchResult, error) {
	return s.repo.StudySearch(query, limit, offset)
}

func (s *SearchService) PerformBlogSearch(query string, limit, offset int) ([]models.BlogSearchResult, error) {
	return s.repo.BlogSearch(query, limit, offset)
}

func (s *SearchService) PerformResumeSearch(query string, limit, offset int) ([]models.ResumeSearchResult, error) {
	return s.repo.ResumeSearch(query, limit, offset)
}

func (s *SearchService) PerformSessionSearch(query string, limit, offset int) ([]models.SessionSearchResult, error) {
	return s.repo.SessionSearch(query, limit, offset)
}

func (s *SearchService) PerformEventSearch(query string, limit, offset int) ([]models.EventSearchResult, error) {
	return s.repo.EventSearch(query, limit, offset)
}

// func (s *SearchService) PerformFinalSearch(index, query string, limit, offset int) ([]models.SearchResult, error) {
// 	return s.repo.FinalSearch(index, query, limit, offset)
// }
