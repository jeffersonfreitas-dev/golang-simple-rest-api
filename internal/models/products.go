package models

import "github.com/google/uuid"

type Product struct {
	ID    uuid.UUID
	Name  string
	Type  string
	Price float64
}

type ProductRequest struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type ProductResponse struct {
	NewProductID uuid.UUID `json:"product_id"`
}
