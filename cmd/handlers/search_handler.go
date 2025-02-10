package handlers

import (
	"net/http"
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
// @Failure 500 {object} map[string]interface{}
// @Router /search/combined [get]
func SearchHandler(c *gin.Context, counter *prometheus.CounterVec, histogram *prometheus.HistogramVec) {
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

	// Perform search
	results, err := service.PerformSearch(req.Index, req.Query, req.Limit, req.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": results})
	counter.WithLabelValues("success").Inc()
	histogram.WithLabelValues("success").Observe(time.Since(startTime).Seconds())
}
