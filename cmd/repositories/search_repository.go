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
		school, _ := hitMap["school"].(string)
		email, _ := hitMap["email"].(string)
		year, _ := hitMap["year"].(int)
		grade, _ := hitMap["grade"].(string)
		stack, _ := hitMap["stack"].([]interface{})
		profileImage, _ := hitMap["profileImage"].(string)
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.UserSearchResult{
			ID:           id,
			Name:         name,
			School:       school,
			Email:        email,
			Year:         year,
			Grade:        grade,
			Stack:        stack,
			ProfileImage: profileImage,
			Index:        "user",
			Score:        score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) ProjectSearch(query string, limit, offset int) ([]models.ProjectSearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index("project").Search(query, &meilisearch.SearchRequest{
		Limit:                int64(limit),
		Offset:               int64(offset),
		ShowRankingScore:     true,
		AttributesToSearchOn: []string{"name", "title", "projectExplain", "teamStacks"},
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.ProjectSearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		name, _ := hitMap["name"].(string)
		title, _ := hitMap["title"].(string)
		projectExplain, _ := hitMap["projectExplain"].(string)
		resultImages, _ := hitMap["resultImages"].([]interface{})
		teamStacks, _ := hitMap["teamStacks"].([]interface{})
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.ProjectSearchResult{
			ID:             id,
			Name:           name,
			Title:          title,
			ProjectExplain: projectExplain,
			ResultImages:   resultImages,
			TeamStacks:     teamStacks,
			Index:          "project",
			Score:          score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) StudySearch(query string, limit, offset int) ([]models.StudySearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index("study").Search(query, &meilisearch.SearchRequest{
		Limit:                int64(limit),
		Offset:               int64(offset),
		ShowRankingScore:     true,
		AttributesToSearchOn: []string{"name", "title", "projectExplain"},
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.StudySearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		name, _ := hitMap["name"].(string)
		title, _ := hitMap["title"].(string)
		projectExplain, _ := hitMap["projectExplain"].(string)
		resultImages, _ := hitMap["resultImages"].([]interface{})
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.StudySearchResult{
			ID:             id,
			Name:           name,
			Title:          title,
			ProjectExplain: projectExplain,
			ResultImages:   resultImages,
			Index:          "study",
			Score:          score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) BlogSearch(query string, limit, offset int) ([]models.BlogSearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index("blog").Search(query, &meilisearch.SearchRequest{
		Limit:                int64(limit),
		Offset:               int64(offset),
		ShowRankingScore:     true,
		AttributesToSearchOn: []string{"title", "userName", "stack"},
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.BlogSearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		title, _ := hitMap["title"].(string)
		url, _ := hitMap["url"].(string)
		date, _ := hitMap["date"].(string)
		userID, _ := hitMap["userID"].(string)
		username, _ := hitMap["userName"].(string)
		userProfileImage, _ := hitMap["userProfileImage"].(string)
		thumbnail, _ := hitMap["thumbnail"].(string)
		stack, ok := hitMap["stack"].([]interface{})
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.BlogSearchResult{
			ID:               id,
			Title:            title,
			URL:              url,
			Date:             date,
			UserID:           userID,
			UserName:         username,
			UserProfileImage: userProfileImage,
			Thumbnail:        thumbnail,
			Stack:            stack,
			Index:            "blog",
			Score:            score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) ResumeSearch(query string, limit, offset int) ([]models.ResumeSearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index("resume").Search(query, &meilisearch.SearchRequest{
		Limit:                int64(limit),
		Offset:               int64(offset),
		ShowRankingScore:     true,
		AttributesToSearchOn: []string{"name", "title", "userName"},
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.ResumeSearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		title, _ := hitMap["title"].(string)
		url, _ := hitMap["url"].(string)
		createdAt, _ := hitMap["createdAt"].(string)
		userID, _ := hitMap["userID"].(string)
		userName, _ := hitMap["userName"].(string)
		userProfileImage, _ := hitMap["userProfileImage"].(string)
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.ResumeSearchResult{
			ID:               id,
			Title:            title,
			URL:              url,
			CreatedAt:        createdAt,
			UserID:           userID,
			UserName:         userName,
			UserProfileImage: userProfileImage,
			Index:            "resume",
			Score:            score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) SessionSearch(query string, limit, offset int) ([]models.SessionSearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index("session").Search(query, &meilisearch.SearchRequest{
		Limit:                int64(limit),
		Offset:               int64(offset),
		ShowRankingScore:     true,
		AttributesToSearchOn: []string{"title", "presenter"},
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.SessionSearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		thumbnail, _ := hitMap["thumbnail"].(string)
		title, _ := hitMap["title"].(string)
		presenter, _ := hitMap["presenter"].(string)
		date, _ := hitMap["date"].(string)
		likeCount, _ := hitMap["likeCount"].(int)
		viewCount, _ := hitMap["viewCount"].(int)
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.SessionSearchResult{
			ID:        id,
			Thumbnail: thumbnail,
			Title:     title,
			Presenter: presenter,
			Date:      date,
			LikeCount: likeCount,
			ViewCount: viewCount,
			Index:     "session",
			Score:     score,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *SearchRepository) EventSearch(query string, limit, offset int) ([]models.EventSearchResult, error) {
	searchRes, err := (*config.MeiliClient).Index("event").Search(query, &meilisearch.SearchRequest{
		Limit:                int64(limit),
		Offset:               int64(offset),
		ShowRankingScore:     true,
		AttributesToSearchOn: []string{"title", "category"},
	})
	if err != nil {
		log.Println("Search error:", err)
		return nil, err
	}

	// Parse results
	results := []models.EventSearchResult{}
	for _, hit := range searchRes.Hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := hitMap["id"].(string)
		category, _ := hitMap["category"].(string)
		title, _ := hitMap["title"].(string)
		score, ok := hit.(map[string]interface{})["_rankingScore"].(float64)
		if !ok {
			continue
		}
		result := models.EventSearchResult{
			ID:       id,
			Title:    title,
			Category: category,
			Index:    "event",
			Score:    score,
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
