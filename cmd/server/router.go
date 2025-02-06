package server

import (
	"github.com/Techeer-Hogwarts/search/cmd/handlers"
	docs "github.com/Techeer-Hogwarts/search/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
			searchGroup.GET("/combined", handlers.SearchHandler)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
