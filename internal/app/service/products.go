package service

import (
	"e-commerce/internal/app/models"
	"e-commerce/pkg/ulid"

	"github.com/shopspring/decimal"
)

type CreateProductRequest struct {
	Name        string          `json:"name"`
	Image       string          `json:"image"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
	Status      string          `json:"status"`
	CategoryID  string          `json:"category_id"`
}

func (s *Service) CreateProducts(r *CreateProductRequest) error {
	product := &models.Product{
		ID:          ulid.New(),
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Status:      r.Status,
		CategoryID:  r.CategoryID,
		Image:       r.Image,
	}
	return e(s.Storage.CreateProduct(product))
}
