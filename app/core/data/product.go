package data

import "github.com/pkg/errors"

type Product struct {
	ID string `gorm:"primarykey"`
	Name string
	ProductType ProductType
	Price uint64
	// beverages and food.
	Volume uint64
	VolumeUnits string
}

func(p Product) Validate() error {
	if !p.ProductType.IsValid() {
		return errors.Errorf("Invalid product type %s", p.ProductType)
	}
	if len(p.Name) == 0 {
		return errors.Errorf("Empty name")
	}
	if len(p.VolumeUnits) == 0 {
		return errors.Errorf("Empty volume units")
	}
	return nil
}


type ProductsRepository interface {
	Store(Product) error
	GetByID(string) (Product, error)
	GetByType(ProductType) ([]Product, error)
	GetByTypeAndVolume(productType ProductType, volume uint64) ([]Product, error)
	GetByNameLike(string) ([]Product, error)
	GetByTypeAndNameLike(ProductType, string) ([]Product, error)
	GetByVolumeAndTypeAndNameLike(uint64, ProductType, string) ([]Product, error)
	GetAll() ([]Product, error)
}
