package service

import (
	"e-commerce/internal/app/models"
	"e-commerce/internal/app/storage"
	"e-commerce/pkg/ulid"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (s *Service) CreateCategory(r *CreateCategoryRequest) error {
	c := &models.Category{
		ID:   ulid.New(),
		Name: r.Name,
	}
	return e(s.Storage.CreateCategory(c))
}

func (s *Service) ReadCategories(r *storage.ReadCategoriesRequest) ([]*models.Category, error) {
	if err := validateReadRequest(r.ReadRequest, []string{"id", "name"}); err != nil {
		return nil, err
	}
	categories, err := s.Storage.ReadCategories(r)
	return categories, e(err)
}
