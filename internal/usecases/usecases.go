package usecases

import (
	"simple-rest-api/internal/models"
	"simple-rest-api/internal/repositories"
	"simple-rest-api/internal/usecases/products"
	"simple-rest-api/internal/usecases/users"

	"github.com/google/uuid"
)

type UseCases struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.CreateUserRequest) (uuid.UUID, error)
	}

	Product interface {
		GetAll() []models.Product
		FindById(uuid uuid.UUID) (models.Product, error)
		Add(product models.ProductRequest) models.ProductResponse
	}
}

func New(repo *repositories.Repositories) *UseCases {
	return &UseCases{User: users.New(repo), Product: products.New(repo)}
}
