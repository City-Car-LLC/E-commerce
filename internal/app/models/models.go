package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Shop struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Logo        string     `json:"logo"`
	OpeningTime time.Time  `json:"opening_time"`
	ClosingTime time.Time  `json:"closing_time"`
	PhoneNumber string     `json:"phone_number"`
	Addresses   []*Address `json:"-"`
	Products    []*Product `json:"-"`
}

type Address struct {
	ID          string `json:"id"`
	ShopID      string `json:"shop_id"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
}

type Product struct {
	ID          string          `json:"id"`
	ShopID      string          `json:"shop_id"`
	CategoryID  string          `json:"category_id"`
	Name        string          `json:"name"`
	Image       string          `json:"image"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
	Status      string          `json:"status"`
	Category    ProductCategory `json:"-"`
}

type ProductCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID        string `json:"id"`
	ShopID    string `json:"shop_id"`
	ProductID string `json:"product_id"`
	Quantity  string `json:"quantity"`
	Status    string `json:"status"`
}
