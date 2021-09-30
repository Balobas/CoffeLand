package repo

import (
	"CoffeLand/app/core/data"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Discount struct {
	db gorm.DB
}

func NewDiscountRepo(db gorm.DB) data.DiscountsRepository {
	return &Discount{db:db}
}

func (d Discount) Store(discount data.Discount) error {
	_, err := d.GetByID(discount.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return d.db.Create(&discount).Error
		}
		return errors.WithStack(err)
	}

	return d.db.Save(discount).Error
}

func (d Discount) GetByID(id string) (data.Discount, error) {
	var discount data.Discount
	result := d.db.Where("id = ?", id).Find(&discount)
	if result.Error != nil {
		return data.Discount{}, result.Error
	}
	if result.RowsAffected == 0 {
		return data.Discount{}, gorm.ErrRecordNotFound
	}

	return discount, nil
}

func (d Discount) GetByNameLike(name string) ([]data.Discount, error) {
	return handleErrorsAndReturnDiscounts(d.db.Where("name LIKE ?", name + "%"))
}

func (d Discount) GetActual() ([]data.Discount, error) {
	t := time.Now().Format("2006-01-02 15:04:05")
	return handleErrorsAndReturnDiscounts(d.db.Where("end_date >= ?", t))
}

func (d Discount) GetByProductType(productType data.ProductType) ([]data.Discount, error) {
	return handleErrorsAndReturnDiscounts(d.db.Where("product_type = ?", productType))
}

func handleErrorsAndReturnDiscounts(where *gorm.DB) ([]data.Discount, error) {
	var discounts []data.Discount
	result := where.Find(&discounts)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return []data.Discount{}, gorm.ErrRecordNotFound
	}
	return discounts, nil
}