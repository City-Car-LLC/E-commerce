package handler

import (
	"e-commerce/internal/app/models"
	"e-commerce/internal/app/service"
	"e-commerce/internal/app/storage"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shopspring/decimal"
)

// CreateProducts godoc
// @Security ApiKeyAuth
// @Tags products
// @Accept json
// @Param body multipart/form-data service.CreateProductRequest true "request"
// @Success 200
// @Router /api/categories/products [post]
func (h Handler) CreateProducts(req *http.Request) (resp interface{}, err error) {
	r := new(service.CreateProductRequest)

	fi := models.File{}
	n := req.Form.Get("name")
	// Retrieve the file from form data
	fi.Content, fi.Meta, err = req.FormFile("file")
	if err != nil {
		return
	}
	defer fi.Content.Close()
	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + n
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, fi.Content)
	if err != nil {
		return
	}
	r.Name = n
	r.Image = n + filepath.Ext(fi.Meta.Filename)
	r.CategoryID = req.Form.Get("category_id")
	r.Description = req.Form.Get("description")
	r.Status = req.Form.Get("status")
	r.Price, err = decimal.NewFromString(req.Form.Get("price"))

	if err != nil {
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
