package models

// SearchRequest는 검색 요청을 나타냅니다.
// @Param query query string true "Search query string"
// @Param limit query int false "Number of results to return (default 10)"
// @Param offset query int false "Offset for pagination (default 0)"
type SearchRequest struct {
	Query  string `json:"query" form:"query" binding:"required"`
	Limit  int    `json:"limit,omitempty" form:"limit"`
	Offset int    `json:"offset,omitempty" form:"offset"`
}

// SearchResult는 검색 결과를 나타냅니다.
type SearchResult struct {
	ID    string `json:"id" form:"id" example:"1"`
	Title string `json:"title" form:"title" example:"스웨거로 문서화하는 방법"`
	Body  string `json:"body" form:"body" example:"스웨거로 문서화하는 방법을 알아봅시다."`
}
