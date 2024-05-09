package database

import (
	"github.com/d4niells/api/internal/entity"
	"gorm.io/gorm"
)

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(ID string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(ID string) error
}

type Product struct {
	DB *gorm.DB
}

func NewProduct(DB *gorm.DB) *Product {
	return &Product{DB}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(ID string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).
			Offset((page - 1) * limit).
			Order("created_at " + sort).
			Find(&products).
			Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(ID string) error {
	product, err := p.FindByID(ID)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
