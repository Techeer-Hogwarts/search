package repositories

import (
	"log"

	"github.com/Techeer-Hogwarts/search/cmd/models"
	"github.com/Techeer-Hogwarts/search/config"
	"github.com/meilisearch/meilisearch-go"
)

type SearchRepository struct{}

// NewSearchRepository creates a new repository instance
func NewSearchRepository() Repository {
	return &SearchRepository{}
}

// Search executes a query in Meilisearch
func (r *SearchRepository) Search(query string, limit, offset int) ([]models.SearchResult, error) {
	index := (*config.MeiliClient).Index("your_index_name")

	searchRes, err := index.Search(query, &meilisearch.SearchRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.SearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}

		result := models.SearchResult{
			ID:    hitMap["id"].(string),
			Title: hitMap["title"].(string),
			Body:  hitMap["body"].(string),
		}
		results = append(results, result)
	}

	return results, nil
}
