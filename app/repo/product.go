package repo

import (
	"CoffeLand/app/core/data"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Product struct {
	db gorm.DB
}

func NewProductRepo(db gorm.DB) data.ProductsRepository {
	return &Product{db: db}
}

func(p Product) GetByID(id string) (data.Product, error) {
	var product data.Product
	result := p.db.Where("id = ?", id).Find(&product)
	if result.Error != nil {
		return data.Product{}, result.Error
	}
	if result.RowsAffected == 0 {
		return data.Product{}, gorm.ErrRecordNotFound
	}

	return product, nil
}

func(p Product) Store(product data.Product) error {
	_, err := p.GetByID(product.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return p.db.Create(&product).Error
		}
		return errors.WithStack(err)
	}

	return p.db.Save(product).Error
}

func(p Product) GetByType(productType data.ProductType) ([]data.Product, error) {
	return handleErrorsAndReturnProducts(p.db.Where("product_type=?", fmt.Sprintf("%s", productType)))
}

func(p Product) GetByTypeAndVolume(productType data.ProductType, volume uint64) ([]data.Product, error) {
	return handleErrorsAndReturnProducts(p.db.Where("product_type=? AND volume=?", fmt.Sprintf("%s", productType), volume))
}

func(p Product) GetByNameLike(name string) ([]data.Product, error) {
	return handleErrorsAndReturnProducts(p.db.Where("name LIKE ?", name + "%"))
}

func (p Product) GetByTypeAndNameLike(productType data.ProductType, name string) ([]data.Product, error) {
	return handleErrorsAndReturnProducts(p.db.Where("product_type=? AND name LIKE ?", productType, name + "%"))
}

func (p Product) GetByVolumeAndTypeAndNameLike(volume uint64, productType data.ProductType, name string) ([]data.Product, error) {
	return handleErrorsAndReturnProducts(p.db.Where("volume = ? AND product_type = ? AND name LIKE ?", volume, productType, name + "%"))
}

func (p Product) GetAll() ([]data.Product, error) {
	return handleErrorsAndReturnProducts(&p.db)
}

func handleErrorsAndReturnProducts(where *gorm.DB) ([]data.Product, error) {
	var products []data.Product
	result := where.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return []data.Product{}, gorm.ErrRecordNotFound
	}
	return products, nil
}
