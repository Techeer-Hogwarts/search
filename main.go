package main

import (
	"log"

	"github.com/Techeer-Hogwarts/search/cmd/server"
)

func main() {
	// Start the server
	srv := server.NewServer()
	log.Println("Server is running on :8080")
	log.Fatal(srv.ListenAndServe())
}
