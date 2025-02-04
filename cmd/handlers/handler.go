package handlers

import "github.com/Techeer-Hogwarts/search/cmd/services"

type Handler struct {
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{}
}
