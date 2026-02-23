package products

import (
	"errors"
	"simple-rest-api/internal/models"

	"github.com/google/uuid"
)

type Products struct {
	products []models.Product
}

func New() *Products {
	return &Products{products: make([]models.Product, 0)}
}

func (p Products) GetAll() []models.Product {
	return p.products
}

func (p *Products) Add(product models.Product) {
	p.products = append(p.products, product)
}

func (p Products) FindById(productId uuid.UUID) (models.Product, error) {
	var prod models.Product

	for _, v := range p.products {
		if v.ID == productId {
			prod = v
		}
	}

	if prod.ID == uuid.Nil {
		return models.Product{}, errors.New("Produto não encontrado com id informado")
	}

	return prod, nil
}
