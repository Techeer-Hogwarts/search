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

type IndexSearchResult map[string]interface{}

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
	ID           string        `json:"id" form:"id" example:"1"`
	Name         string        `json:"name" form:"name" example:"윤정은"`
	School       string        `json:"school" form:"school" example:"서울대학교"`
	Email        string        `json:"email" form:"email" example:"test@gmail.com"`
	Year         int           `json:"year" form:"year" example:"7"`
	Grade        string        `json:"grade" form:"grade" example:"3"`
	Stack        []interface{} `json:"stack" form:"stack" example:"Go"`
	ProfileImage string        `json:"profileImage" form:"profileImage" example:"https://example.com/profile.jpg"`
	Index        string        `json:"index" form:"index" example:"user"`
	Score        float64       `json:"score" form:"score" example:"0.99"`
}

type ProjectSearchResult struct {
	ID             string        `json:"id" form:"id" example:"1"`
	Name           string        `json:"name" form:"name" example:"프로젝트 이름"`
	Title          string        `json:"title" form:"title" example:"프로젝트 제목"`
	ProjectExplain string        `json:"projectExplain" form:"projectExplain" example:"프로젝트 설명"`
	ResultImages   []interface{} `json:"resultImages" form:"resultImages" example:"https://example.com/result.jpg"`
	TeamStacks     []interface{} `json:"teamStacks" form:"teamStacks" example:"Go"`
	Index          string        `json:"index" form:"index" example:"user"`
	Score          float64       `json:"score" form:"score" example:"0.99"`
}

type StudySearchResult struct {
	ID             string        `json:"id" form:"id" example:"1"`
	Name           string        `json:"name" form:"name" example:"프로젝트 이름"`
	Title          string        `json:"title" form:"title" example:"프로젝트 제목"`
	ProjectExplain string        `json:"projectExplain" form:"projectExplain" example:"프로젝트 설명"`
	ResultImages   []interface{} `json:"resultImages" form:"resultImages" example:"https://example.com/result.jpg"`
	Index          string        `json:"index" form:"index" example:"user"`
	Score          float64       `json:"score" form:"score" example:"0.99"`
}

type BlogSearchResult struct {
	ID               string        `json:"id" form:"id" example:"1"`
	Title            string        `json:"title" form:"title" example:"블로그 제목"`
	URL              string        `json:"url" form:"url" example:"https://example.com/blog"`
	Date             string        `json:"date" form:"date" example:"2021-01-01"`
	UserID           string        `json:"userID" form:"userID" example:"1"`
	UserName         string        `json:"userName" form:"userName" example:"윤정은"`
	UserProfileImage string        `json:"userProfileImage" form:"userProfileImage" example:"https://example.com/profile.jpg"`
	Thumbnail        string        `json:"thumbnail" form:"thumbnail" example:"https://example.com/thumbnail.jpg"`
	Stack            []interface{} `json:"stack" form:"stack" example:"Go"`
	Index            string        `json:"index" form:"index" example:"user"`
	Score            float64       `json:"score" form:"score" example:"0.99"`
}

type ResumeSearchResult struct {
	ID               string  `json:"id" form:"id" example:"1"`
	Title            string  `json:"title" form:"title" example:"이력서 제목"`
	URL              string  `json:"url" form:"url" example:"https://example.com/resume"`
	CreatedAt        string  `json:"createdAt" form:"createdAt" example:"2021-01-01"`
	UserID           string  `json:"userID" form:"userID" example:"1"`
	UserName         string  `json:"userName" form:"userName" example:"윤정은"`
	UserProfileImage string  `json:"userProfileImage" form:"userProfileImage" example:"https://example.com/profile.jpg"`
	Index            string  `json:"index" form:"index" example:"user"`
	Score            float64 `json:"score" form:"score" example:"0.99"`
}

type SessionSearchResult struct {
	ID        string  `json:"id" form:"id" example:"1"`
	Thumbnail string  `json:"thumbnail" form:"thumbnail" example:"https://example.com/thumbnail.jpg"`
	Title     string  `json:"title" form:"title" example:"세션 제목"`
	Presenter string  `json:"presenter" form:"presenter" example:"윤정은"`
	Date      string  `json:"date" form:"date" example:"2021-01-01"`
	LikeCount int     `json:"likeCount" form:"likeCount" example:"10"`
	ViewCount int     `json:"viewCount" form:"viewCount" example:"100"`
	Index     string  `json:"index" form:"index" example:"user"`
	Score     float64 `json:"score" form:"score" example:"0.99"`
}

type EventSearchResult struct {
	ID       string  `json:"id" form:"id" example:"1"`
	Category string  `json:"category" form:"category" example:"세션"`
	Title    string  `json:"title" form:"title" example:"세션 제목"`
	Index    string  `json:"index" form:"index" example:"user"`
	Score    float64 `json:"score" form:"score" example:"0.99"`
}
