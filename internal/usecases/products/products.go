package products

import (
	"simple-rest-api/internal/models"
	"simple-rest-api/internal/repositories"

	"github.com/google/uuid"
)

type ProductCase struct {
	repo *repositories.Repositories
}

func New(repo *repositories.Repositories) *ProductCase {
	return &ProductCase{repo: repo}
}

func (p ProductCase) GetAll() []models.Product {
	return p.repo.Product.GetAll()
}

func (p ProductCase) FindById(uuid uuid.UUID) (models.Product, error) {
	product, err := p.repo.Product.FindById(uuid)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p *ProductCase) Add(product models.ProductRequest) models.ProductResponse {

	//todo: Validar se nome já existe

	prod := models.Product{
		ID:    uuid.New(),
		Name:  product.Name,
		Type:  product.Type,
		Price: product.Price,
	}

	p.repo.Product.Add(prod)
	return models.ProductResponse{NewProductID: prod.ID}
}
