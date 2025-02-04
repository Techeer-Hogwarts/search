package services

import "github.com/Techeer-Hogwarts/search/cmd/repositories"

type Service struct {
}

// NewService creates a new instance of Service with all required services.
func NewService(repo *repositories.Repository) *Service {
	return &Service{}
}
