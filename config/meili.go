package config

import (
	"log"

	"github.com/meilisearch/meilisearch-go"
)

var MeiliClient *meilisearch.ServiceManager

// Initialize Meilisearch client
func InitMeilisearch() {
	clientAddr := GetEnvVarAsString("MEILISEARCH_ADDR", "http://localhost:7700")
	clientAPIKey := GetEnvVarAsString("MEILISEARCH_API_KEY", "")
	client := meilisearch.New(clientAddr, meilisearch.WithAPIKey(clientAPIKey))
	MeiliClient = &client
	log.Println("Meilisearch client initialized")
}
