package service

import (
	"e-commerce/internal/app/models"
	"e-commerce/pkg/ulid"
	"time"
)

type CreateShopRequest struct {
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`
	PhoneNumber string    `json:"phone_number"`
}

func (s *Service) CreateShop(r *CreateShopRequest) error {
	c := &models.Shop{
		ID:          ulid.New(),
		Name:        r.Name,
		Logo:        r.Logo,
		PhoneNumber: r.PhoneNumber,
		OpeningTime: r.OpeningTime,
		ClosingTime: r.ClosingTime,
	}

	return e(s.Storage.CreateShop(c))
}
