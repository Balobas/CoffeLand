package usecases

import (
	"CoffeLand/app/core/data"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ProductUsecases struct {
	data.Core
}

func NewProductUsecases(core data.Core) *ProductUsecases {
	return &ProductUsecases{core}
}

func(p *ProductUsecases) GetByFilters(productType data.ProductType, volume uint64, name string) ([]data.Product, error) {
	if len(productType) == 0 && volume == 0 && len(name) == 0 {
		return p.Core.ProductRepo.GetAll()
	}
	if len(name) != 0 && volume == 0 && len(productType) == 0 {
		return p.Core.ProductRepo.GetByNameLike(name)
	}
	if len(name) == 0 && volume == 0 {
		return p.Core.ProductRepo.GetByType(productType)
	}
	if len(productType) != 0 && len(name) != 0 && volume == 0 {
		return p.Core.ProductRepo.GetByTypeAndNameLike(productType, name)
	}

	if len(name) == 0 && len(productType) != 0 {
		return p.Core.ProductRepo.GetByTypeAndVolume(productType, volume)
	}
	if len(productType) == 0 {
		return nil, errors.New("Invalid filters combination: volume can be used only with type")
	}

	return p.Core.ProductRepo.GetByVolumeAndTypeAndNameLike(volume, productType, name)
}

func(p *ProductUsecases) GetAll() ([]data.Product, error) {
	return p.Core.ProductRepo.GetAll()
}

func(p *ProductUsecases) putProduct(product data.Product) (string, error) {
	if err := product.Validate(); err != nil {
		return "", errors.WithStack(err)
	}

	if len(product.ID) == 0 {
		product.ID = uuid.New().String()
	}

	if err := p.Core.ProductRepo.Store(product); err != nil {
		return product.ID, errors.WithStack(err)
	}

	return product.ID, nil
}