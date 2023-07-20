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
// @Param request body service.CreateCategoryRequest true "request" Format(json)
// @Success 200
// @Router /api/categories [post]
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
// @Param  request query storage.ReadCategoriesRequest true "request" Format(json)
// @Success 200 {object} []models.Category
// @Router /api/categories [get]
func (h Handler) ReadCategories(req *http.Request) (resp interface{}, err error) {
	r := new(storage.ReadCategoriesRequest)
	if err = unmarshal(req, r); err != nil {
		return
	}
	resp, err = h.Service.ReadCategories(r)
	return
}
