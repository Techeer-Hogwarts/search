package handlers

import (
	"net/http"
	"sort"
	"time"

	"github.com/Techeer-Hogwarts/search/cmd/models"
	"github.com/Techeer-Hogwarts/search/cmd/repositories"
	"github.com/Techeer-Hogwarts/search/cmd/services"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// SearchHandler handles search requests
// @Summary Search in Meilisearch
// @Description Query Meilisearch and return results
// @Tags search
// @Accept json
// @Produce json
// @Param index query string true "name of index"
// @Param query query string true "Search query string"
// @Param limit query int false "Number of results to return (default 10)"
// @Param offset query int false "Offset for pagination (default 0)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /search [get]
func SearchHandler(c *gin.Context, counter *prometheus.CounterVec, histogram *prometheus.HistogramVec) {

	// contains checks if a slice contains a specific element
	startTime := time.Now()
	var req models.SearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dependency injection
	repo := repositories.NewSearchRepository()
	service := services.NewSearchService(repo)
	// index := "documents"
	validJWT, _ := c.Get("valid_jwt")
	var allowedIndex []string
	if validJWT == true {
		allowedIndex = []string{"user", "resume", "blog", "session", "projectTeam", "studyTeam", "event"}
	} else {
		allowedIndex = []string{"blog", "users", "event"}
	}
	// available indexes: user, resume, blog, session, projectTeam, studyTeam, event
	// Perform search
	if !contains(allowedIndex, req.Index) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Index Not Allowed"})
		return
	}
	if req.Index == "user" {
		results, err := service.PerformUserSearch(req.Index, req.Query, req.Limit, req.Offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"results": results})
		counter.WithLabelValues("success").Inc()
		histogram.WithLabelValues("success").Observe(time.Since(startTime).Seconds())
		return
	}
	results, err := service.PerformSearch(req.Index, req.Query, req.Limit, req.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": results})
	counter.WithLabelValues("success").Inc()
	histogram.WithLabelValues("success").Observe(time.Since(startTime).Seconds())
}

// BasicSearchHandler handles search requests
// @Summary Basic Combined Search in Meilisearch
// @Description Query All Index in Meilisearch and return results
// @Tags search
// @Accept json
// @Produce json
// @Param query query string true "Search query string"
// @Param limit query int false "Number of results to return (default 10)"
// @Param offset query int false "Offset for pagination (default 0)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /search/basic [get]
func BasicSearchHandler(c *gin.Context, counter *prometheus.CounterVec, histogram *prometheus.HistogramVec) {
	startTime := time.Now()
	var req models.CombinedSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dependency injection
	repo := repositories.NewSearchRepository()
	service := services.NewSearchService(repo)
	// available indexes: user, resume, blog, session, projectTeam, studyTeam, event
	validated := []string{"resume", "blog", "session", "projectTeam", "studyTeam", "event"}
	invalidated := []string{"blog", "event"}

	var results []models.CombinedSearchResult
	validJWT, _ := c.Get("valid_jwt")
	if validJWT == false {
		for _, index := range invalidated {
			// Perform search
			result, err := service.PerformBasicSearch(index, req.Query, req.Limit, req.Offset)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
				return
			}
			results = append(results, result...)
		}
	} else {
		for _, index := range validated {
			// Perform search
			result, err := service.PerformBasicSearch(index, req.Query, req.Limit, req.Offset)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
				return
			}
			results = append(results, result...)
		}
	}

	// Sort results by score
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	c.JSON(http.StatusOK, gin.H{"results": results})
	counter.WithLabelValues("success").Inc()
	histogram.WithLabelValues("success").Observe(time.Since(startTime).Seconds())
}

// FinalSearchHandler handles search requests
// @Summary Final Combined Search in Meilisearch
// @Description Query All Index in Meilisearch and return results
// @Tags search
// @Accept json
// @Produce json
// @Param query query string true "Search query string"
// @Param limit query int false "Number of results to return (default 10)"
// @Param offset query int false "Offset for pagination (default 0)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /search/final [get]
func FinalSearchHandler(c *gin.Context, counter *prometheus.CounterVec, histogram *prometheus.HistogramVec) {
	startTime := time.Now()
	var req models.CombinedSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dependency injection
	repo := repositories.NewSearchRepository()
	service := services.NewSearchService(repo)
	// available indexes: user, resume, blog, session, projectTeam, studyTeam, event
	validated := []string{"resume", "blog", "session", "projectTeam", "studyTeam", "event"}
	invalidated := []string{"blog", "event"}
	var finalResult models.FinalSearchResult
	finalResult.Result = make(models.IndexSearchResult)
	var results []models.CombinedSearchResult
	validJWT, _ := c.Get("valid_jwt")
	if validJWT == false {
		for _, index := range invalidated {
			// Perform search
			result, err := service.PerformFinalSearch(index, req.Query, req.Limit, req.Offset)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
				return
			}
			sort.Slice(results, func(i, j int) bool {
				return results[i].Score > results[j].Score
			})
			finalResult.Result[index] = result
		}
	} else {
		for _, index := range validated {
			// Perform search
			result, err := service.PerformFinalSearch(index, req.Query, req.Limit, req.Offset)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
				return
			}
			sort.Slice(results, func(i, j int) bool {
				return results[i].Score > results[j].Score
			})
			finalResult.Result[index] = result
		}
	}

	// Sort results by score
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	c.JSON(http.StatusOK, finalResult)
	counter.WithLabelValues("success").Inc()
	histogram.WithLabelValues("success").Observe(time.Since(startTime).Seconds())
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
