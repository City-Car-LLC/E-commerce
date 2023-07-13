package storage

import (
	"e-commerce/internal/app/models"
	"errors"

	"gorm.io/gorm"
)

type Storage struct {
	ORM *gorm.DB
}

func (s Storage) NotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (s Storage) CreateShop(m *models.Shop) error {
	return s.ORM.Create(m).Error
}

func (s Storage) CreateCategory(m *models.Category) error {
	return s.ORM.Create(m).Error
}

func (s Storage) CreateOrder(m *models.Order) error {
	return s.ORM.Create(m).Error
}

func (s Storage) CreateAddress(m *models.Address) error {
	return s.ORM.Create(m).Error
}

func (s Storage) ReadOrder(id string) (o *models.Order, err error) {
	err = s.ORM.Take(&o, "id", id).Error
	return
}

func (s Storage) ReadShop(id string) (sh *models.Shop, err error) {
	err = s.ORM.Take(&sh, "id", id).Error
	return
}

func (s Storage) ReadCategory(id string) (c *models.Category, err error) {
	err = s.ORM.Take(&c, "id", id).Error
	return
}

func (s Storage) UpdateCategory(c *models.Category) error {
	return s.ORM.Save(c).Error
}

func (s Storage) UpdateShop(c *models.Shop) error {
	return s.ORM.Save(c).Error
}

func (s Storage) UpdateAddress(m *models.Address) error {
	return s.ORM.Save(m).Error
}

func (s Storage) UpdateProduct(m *models.Product) error {
	return s.ORM.Save(m).Error
}

type ReadCategoriesRequest struct {
	*ReadRequest
}

func (s Storage) ReadCategories(r *ReadCategoriesRequest) (ms []*models.Category, err error) {
	err = newFilterQuery(s.ORM, r.ReadRequest).q.Find(&ms).Error
	return
}
