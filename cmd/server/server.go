package server

import (
	"fmt"
	"log"

	"github.com/Techeer-Hogwarts/search/cmd/database"
	"github.com/Techeer-Hogwarts/search/cmd/handlers"
	"github.com/Techeer-Hogwarts/search/cmd/repositories"
	"github.com/Techeer-Hogwarts/search/cmd/services"
)

func Start(port string) {
	db := database.SetupDatabase()
	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	router := setupRouter(handler)
	router.Run(fmt.Sprintf(":%s", port))
	log.Printf("Server started on port %s", port)
}
