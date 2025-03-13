package repositories

import "github.com/Techeer-Hogwarts/search/cmd/models"

type Repository interface {
	Search(index, query string, limit, offset int) ([]models.CombinedSearchResult, error)
	TitleSearch(index, query string, limit, offset int) ([]models.CombinedSearchResult, error)
	UserSearch(query string, limit, offset int) ([]models.UserSearchResult, error)
	ProjectSearch(query string, limit, offset int) ([]models.ProjectSearchResult, error)
	StudySearch(query string, limit, offset int) ([]models.StudySearchResult, error)
	BlogSearch(query string, limit, offset int) ([]models.BlogSearchResult, error)
	ResumeSearch(query string, limit, offset int) ([]models.ResumeSearchResult, error)
	SessionSearch(query string, limit, offset int) ([]models.SessionSearchResult, error)
	EventSearch(query string, limit, offset int) ([]models.EventSearchResult, error)
	StackSearch(query string, limit, offset int) ([]models.StackSearchResult, error)
	// FinalSearch(index, query string, limit, offset int) ([]models.SearchResult, error)
}
