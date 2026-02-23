package main

import (
	"simple-rest-api/internal/handlers"
	"simple-rest-api/internal/repositories"
	"simple-rest-api/internal/usecases/products"
	"simple-rest-api/internal/usecases/users"
)

// handler <- usecases <- repositories
func main() {

	repo := repositories.New()
	useCases := users.New(repo)
	productCases := products.New(repo)
	h := handlers.New(useCases, productCases)

	h.Listen(8090)
}
