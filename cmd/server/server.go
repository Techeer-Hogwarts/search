package server

import (
	"log"
	"net/http"

	"github.com/Techeer-Hogwarts/search/config"
)

func NewServer() *http.Server {
	r := setupRouter()
	config.InitMeilisearch()

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Server initialized")
	return server
}
