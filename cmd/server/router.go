package server

import (
	"github.com/Techeer-Hogwarts/search/cmd/handlers"
	docs "github.com/Techeer-Hogwarts/search/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	indexAccessCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "search_access_total",
			Help: "Total number of times search handler was accessed",
		},
		[]string{"status"}, // Success or failure status
	)

	searchDurationHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "search_processing_duration_seconds",
			Help:    "Time taken to process search requests",
			Buckets: prometheus.DefBuckets, // Default Prometheus buckets
		},
		[]string{"status"},
	)
)

// func init() {
// 	prometheus.MustRegister(indexAccessCounter)
// 	prometheus.MustRegister(searchDurationHistogram)
// 	prometheus.MustRegister(collectors.NewGoCollector())
// }

// setupRouter sets up the routes for the application.
func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(jsonLoggerMiddleware()) // logging
	// CORS middleware 다 허용
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	docs.SwaggerInfo.Title = "Techeerzip Search API"
	docs.SwaggerInfo.Description = "Search Engine API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	apiGroup := router.Group("/api/v1")
	{
		// search routes
		searchGroup := apiGroup.Group("/search")
		{
			searchGroup.GET("/combined", func(c *gin.Context) {
				handlers.SearchHandler(c, indexAccessCounter, searchDurationHistogram)
			})
		}

	}
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
