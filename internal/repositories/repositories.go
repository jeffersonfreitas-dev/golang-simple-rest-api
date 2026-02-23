package repositories

import (
	"simple-rest-api/internal/models"
	"simple-rest-api/internal/repositories/products"
	"simple-rest-api/internal/repositories/users"

	"github.com/google/uuid"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailExists(email string) bool
	}
	Product interface {
		GetAll() []models.Product
		Add(product models.Product)
		FindById(id uuid.UUID) (models.Product, error)
	}
}

func New() *Repositories {
	return &Repositories{User: users.New(), Product: products.New()}
}
