package repositories

import "github.com/Techeer-Hogwarts/search/cmd/models"

type Repository interface {
	Search(index, query string, limit, offset int) ([]models.CombinedSearchResult, error)
	TitleSearch(index, query string, limit, offset int) ([]models.CombinedSearchResult, error)
	UserSearch(query string, limit, offset int) ([]models.UserSearchResult, error)
	FinalSearch(index, query string, limit, offset int) ([]models.SearchResult, error)
}
