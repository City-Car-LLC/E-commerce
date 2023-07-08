package entities

import "time"

// ShopType представляет тип магазина
type ShopType string

const (
	ShopTypeAutoParts   ShopType = "AutoParts"   // Магазин автозапчастей
	ShopTypeAutomobiles ShopType = "Automobiles" // Магазин автомобилей
)

// Location представляет местоположение магазина
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
}

// LegalEntityType представляет тип юридического лица
type LegalEntityType string

const (
	LegalEntityTypeIndividual LegalEntityType = "Individual" // ИП
	LegalEntityTypeLegal      LegalEntityType = "Legal"      // Юридическое лицо
)

// Shop представляет сущность магазина
type Shop struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	LogoURL         string          `json:"logo_url"`
	WorkingHours    TimeRange       `json:"working_hours"`
	PhoneNumber     string          `json:"phone_number"`
	Type            ShopType        `json:"type"`         // Тип магазина (автозапчасти или автомобили)
	LegalEntityType LegalEntityType `json:"legal_entity"` // Тип организации (ИП или юридическое лицо)
	Parts           []Product       `json:"parts"`        // Автозапчасти, принадлежащие данному магазину
	Location        Location        `json:"location"`
}

// TimeRange представляет временной диапазон
type TimeRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url"`
	Price       float64  `json:"price"`
	Status      string   `json:"status"`
	Category    Category `json:"-"`
	ShopID      string   `json:"shop_id"` // Идентификатор магазина, к которому принадлежит товар
}

// Category представляет сущность категории товара
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
