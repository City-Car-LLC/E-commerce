package handler

import (
	"e-commerce/internal/app/models"
	"e-commerce/internal/app/service"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// CreateShops godoc
// @Security ApiKeyAuth
// @Tags shops
// @Accept json
// @Param body multipart/form-data service.CreateShopRequest true "request"
// @Success 200
// @Router /api/shops/categories/products [post]
func (h Handler) CreateShops(req *http.Request) (resp interface{}, err error) {
	r := new(service.CreateShopRequest)

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
	r.Logo = n + filepath.Ext(fi.Meta.Filename)
	r.PhoneNumber = req.Form.Get("phone_number")
	openTime := req.Form.Get("opening_time")
	closingTime := req.Form.Get("closing_time")
	layout := "2006-01-02 15:04:05"
	r.OpeningTime, err = time.Parse(layout, openTime)
	if err != nil {
		return
	}
	r.ClosingTime, err = time.Parse(layout, closingTime)
	if err != nil {
		log.Println("Error on Parsing ClosingTime")
		return
	}
	err = h.Service.CreateShop(r)
	return
}
