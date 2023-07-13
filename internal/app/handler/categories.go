package handler

import (
	"e-commerce/internal/app/service"
	"e-commerce/internal/app/storage"
	"net/http"
)

// CreateCategories godoc
// @Security ApiKeyAuth
// @Tags categories
// @Accept json
// @Param body service.CreateCategoryRequest true "request"
// @Success 200
// @Router /api/regions/categories [post]
func (h Handler) CreateCategories(req *http.Request) (resp interface{}, err error) {
	r := new(service.CreateCategoryRequest)
	if err = unmarshal(req, r); err != nil {
		return
	}
	err = h.Service.CreateCategory(r)
	return
}

// ReadCategories godoc
// @Security ApiKeyAuth
// @Tags cities
// @Produce json
// @Param query storage.ReadCategoriesRequest true "request"
// @Success 200 {object} []models.Category
// @Router /api/regions/categories [get]
func (h Handler) ReadCategories(req *http.Request) (resp interface{}, err error) {
	r := new(storage.ReadCategoriesRequest)
	if err = unmarshal(req, r); err != nil {
		return
	}
	resp, err = h.Service.ReadCategories(r)
	return
}
