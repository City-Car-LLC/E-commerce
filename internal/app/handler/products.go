package handler

import (
	"e-commerce/internal/app/service"
	"e-commerce/internal/app/storage"
	"net/http"
)

// CreateProducts godoc
// @Security ApiKeyAuth
// @Tags products
// @Accept json
// @Param body service.CreateProductRequest true "request"
// @Success 200
// @Router /api/categories/products [post]
func (h Handler) CreateProducts(req *http.Request) (resp interface{}, err error) {
	r := new(service.CreateProductRequest)
	if err = unmarshal(req, r); err != nil {
		return
	}
	err = h.Service.CreateProducts(r)
	return
}

// ReadProducts godoc
// @Security ApiKeyAuth
// @Tags products
// @Produce json
// @Param query storage.ReadProductsRequest true "request"
// @Success 200 {object} []models.Product
// @Router /api/categories/products [get]
func (h Handler) ReadProducts(req *http.Request) (resp interface{}, err error) {
	r := new(storage.ReadProductsRequest)
	if err = unmarshal(req, r); err != nil {
		return
	}
	resp, err = h.Service.Storage.ReadProducts(r)
	return
}
