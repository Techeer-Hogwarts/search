package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Techeer-Hogwarts/search/config"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	// Wait for Meilisearch to be ready
	time.Sleep(2 * time.Second)
	// config.LoadEnvFile("../.env")

	meiliHost := config.GetEnvVarAsString("MEILISEARCH_ADDR", "http://localhost:7700")
	meiliKey := config.GetEnvVarAsString("MEILISEARCH_API_KEY", "masterKey")

	client := meilisearch.New(meiliHost, meilisearch.WithAPIKey(meiliKey))

	// Define index settings
	index := client.Index("test")

	// Sample data
	documents := []map[string]interface{}{
		{"id": "1", "title": "Go Programming", "body": "Learn Go for backend development"},
		{"id": "2", "title": "Docker Guide", "body": "Understanding containers with Docker"},
		{"id": "3", "title": "Meilisearch Introduction", "body": "Fast search engine for applications"},
		{"id": "4", "title": "Kubernetes Basics", "body": "Orchestrate containers with Kubernetes"},
		{"id": "5", "title": "ReactJS Tutorial", "body": "Build interactive UIs with React"},
		{"id": "6", "title": "VueJS Fundamentals", "body": "Learn VueJS for frontend development"},
		{"id": "7", "title": "NodeJS Essentials", "body": "Introduction to NodeJS for backend"},
		{"id": "8", "title": "Python Crash Course Docker", "body": "Python programming for beginners"},
		{"id": "9", "title": "Rust Programming", "body": "Do systems programming with Rust"},
		{"id": "10", "title": "백한결과 친구들", "body": "백한결 정유진 라이언 레츠고"},
		{"id": "11", "title": "현진 언니 취업", "body": "현진 언니 취업 축하합니다"},
	}

	// Insert data
	_, err := index.AddDocuments(documents)
	if err != nil {
		log.Fatalf("Error inserting documents: %v", err)
	}

	fmt.Println("Sample data successfully inserted into Meilisearch!")
}
