package main

import (
	"simple-rest-api/internal/handlers"
	"simple-rest-api/internal/repositories"
	"simple-rest-api/internal/usecases"
)

// handler <- usecases <- repositories
func main() {

	repo := repositories.New()
	useCases := usecases.New(repo)
	h := handlers.New(useCases)

	h.Listen(8090)
}
