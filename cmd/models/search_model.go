package models

// SearchRequest는 검색 요청을 나타냅니다.
type SearchRequest struct {
	Index  string `json:"index" form:"index" binding:"required"`
	Query  string `json:"query" form:"query" binding:"required"`
	Limit  int    `json:"limit,omitempty" form:"limit"`
	Offset int    `json:"offset,omitempty" form:"offset"`
}

// CombinedSearchRequest는 통합 검색 요청을 나타냅니다.
type CombinedSearchRequest struct {
	Query  string `json:"query" form:"query" binding:"required"`
	Limit  int    `json:"limit,omitempty" form:"limit"`
	Offset int    `json:"offset,omitempty" form:"offset"`
}

// SearchResult는 검색 결과를 나타냅니다.
type SearchResult struct {
	ID    string  `json:"id" form:"id" example:"1"`
	Index string  `json:"index" form:"index" example:"blog"`
	Score float64 `json:"score" form:"score" example:"0.99"`
}

type IndexSearchResult map[string][]SearchResult

type FinalSearchResult struct {
	Result IndexSearchResult `json:"result" form:"result"`
}

// CombinedSearchResult는 검색 결과를 나타냅니다.
type CombinedSearchResult struct {
	ID    string  `json:"id" form:"id" example:"1"`
	Title string  `json:"title" form:"title" example:"스웨거로 문서화하는 방법"`
	Index string  `json:"index" form:"index" example:"blog"`
	Score float64 `json:"score" form:"score" example:"0.99"`
}

type UserSearchResult struct {
	ID           string  `json:"id" form:"id" example:"1"`
	Name         string  `json:"name" form:"name" example:"윤정은"`
	Email        string  `json:"email" form:"email" example:"test@gmail.com"`
	Year         int     `json:"year" form:"year" example:"7"`
	ProfileImage string  `json:"profileImage" form:"profileImage" example:"https://example.com/profile.jpg"`
	Index        string  `json:"index" form:"index" example:"user"`
	Score        float64 `json:"score" form:"score" example:"0.99"`
}
