package main

import (
	"github.com/Techeer-Hogwarts/search/cmd/server"
	"github.com/Techeer-Hogwarts/search/config"
)

func main() {
	port := config.GetEnvVarAsString("PORT", "8080")
	server.Start(port)
}
