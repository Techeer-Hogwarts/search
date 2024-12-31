package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Basic Hello World on port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	log.Println("Server started on: http://localhost:8085")
	http.ListenAndServe(":8085", nil)
}
