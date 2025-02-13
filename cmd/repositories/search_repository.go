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
func (r *SearchRepository) Search(index, query string, limit, offset int) ([]models.CombinedSearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index(index).Search(query, &meilisearch.SearchRequest{
		Limit:            int64(limit),
		Offset:           int64(offset),
		ShowRankingScore: true,
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.CombinedSearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		title, _ := hitMap["title"].(string)
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.CombinedSearchResult{
			ID:    id,
			Title: title,
			Index: index,
			Score: score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) TitleSearch(index, query string, limit, offset int) ([]models.CombinedSearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index(index).Search(query, &meilisearch.SearchRequest{
		Limit:                int64(limit),
		Offset:               int64(offset),
		ShowRankingScore:     true,
		AttributesToSearchOn: []string{"title"},
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.CombinedSearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		title, _ := hitMap["title"].(string)
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.CombinedSearchResult{
			ID:    id,
			Title: title,
			Index: index,
			Score: score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) UserSearch(query string, limit, offset int) ([]models.UserSearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index("user").Search(query, &meilisearch.SearchRequest{
		Limit:                int64(limit),
		Offset:               int64(offset),
		ShowRankingScore:     true,
		AttributesToSearchOn: []string{"name"},
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.UserSearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		name, _ := hitMap["name"].(string)
		email, _ := hitMap["email"].(string)
		year, _ := hitMap["year"].(int)
		profileImage, _ := hitMap["profileImage"].(string)
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.UserSearchResult{
			ID:           id,
			Name:         name,
			Email:        email,
			Year:         year,
			ProfileImage: profileImage,
			Index:        "user",
			Score:        score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) FinalSearch(index, query string, limit, offset int) ([]models.SearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index(index).Search(query, &meilisearch.SearchRequest{
		Limit:            int64(limit),
		Offset:           int64(offset),
		ShowRankingScore: true,
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
		id, _ := hitMap["id"].(string)
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.SearchResult{
			ID:    id,
			Index: index,
			Score: score,
		}
		results = append(results, result)
	}

	return results, nil
}
