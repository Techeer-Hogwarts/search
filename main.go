package main

import (
	"log"

	"github.com/Techeer-Hogwarts/search/cmd/server"
)

// @securityDefinitions.apikey cookie
// @in cookie
// @name access_token
func main() {
	srv := server.NewServer()
	log.Println("Server is running on :8080")
	log.Fatal(srv.ListenAndServe())
}
