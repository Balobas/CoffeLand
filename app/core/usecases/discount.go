package usecases

import (
	"CoffeLand/app/core/data"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type DiscountUsecases struct {
	*data.Core
}

func NewDiscountUsecases(core *data.Core) *DiscountUsecases {
	return &DiscountUsecases{core}
}

func(d *DiscountUsecases) GetByNameLike(name string) ([]data.Discount, error) {
	return d.Core.DiscountRepo.GetByNameLike(name)
}

func(d *DiscountUsecases) GetActual() ([]data.Discount, error) {
	return d.Core.DiscountRepo.GetActual()
}

func(d *DiscountUsecases) GetByProductType(productType data.ProductType) ([]data.Discount, error) {
	return d.Core.DiscountRepo.GetByProductType(productType)
}

func(d *DiscountUsecases) putDiscount(discount data.Discount) (string, error) {
	if err := discount.Validate(); err != nil {
		return "", errors.WithStack(err)
	}
	if len(discount.ID) == 0 {
		discount.ID = uuid.New().String()
	}
	if err := d.Core.DiscountRepo.Store(discount); err != nil {
		return discount.ID, errors.WithStack(err)
	}
	return discount.ID, nil
}