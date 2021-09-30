package data

import "github.com/pkg/errors"

type Discount struct {
	ID string `gorm:"primarykey"`
	Name string
	StartDate string
	EndDate string
	ProductType ProductType
	Description string
}

func(d Discount) Validate() error {
	if !d.ProductType.IsValid() {
		return errors.Errorf("Invalid product type %s", d.ProductType)
	}
	if len(d.Name) == 0 {
		return errors.Errorf("Empty name")
	}
	if len(d.Description) == 0 {
		return errors.Errorf("Empty name")
	}
	return nil
}

type DiscountsRepository interface {
	Store(Discount) error
	GetByID(string) (Discount, error)
	GetByNameLike(string) ([]Discount, error)
	GetActual() ([]Discount, error)
	GetByProductType(ProductType) ([]Discount, error)
}
