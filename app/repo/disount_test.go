package repo

import (
	"CoffeLand/app/core/data"
	"CoffeLand/app/interfaces/database"
	"testing"
)

func getDiscountModel(t *testing.T) data.DiscountsRepository {
	db, err := database.MySQLDB()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return NewDiscountRepo(db)
}

func TestDiscountModel_Store(t *testing.T) {
	dm := getDiscountModel(t)

	err := dm.Store(data.Discount{
		ID:          "kek",
		Name:        "2 po zene 1",
		StartDate:   "2021-09-12",
		EndDate:     "2021-09-25",
		ProductType: "coffee",
		Description: "akzia",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestDiscountModel_GetByID(t *testing.T) {
	discount, err := getDiscountModel(t).GetByID("kek")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(discount)

	_, err = getDiscountModel(t).GetByID("adasdad")
	if err == nil {
		t.Error("expected not found records error")
		t.FailNow()
	}
}

func TestDiscountModel_GetByProductType(t *testing.T) {
	discounts, err := getDiscountModel(t).GetByProductType("coffee")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(discounts)

	_, err = getDiscountModel(t).GetByProductType("dsofndonfdaod")
	if err == nil {
		t.Error("expected not found records error")
		t.FailNow()
	}
}

func TestDiscountModel_GetByNameLike(t *testing.T) {
	discounts, err := getDiscountModel(t).GetByNameLike("2 po")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(discounts)

	_, err = getDiscountModel(t).GetByNameLike("asads")
	if err == nil {
		t.Error("expected not found records error")
		t.FailNow()
	}
}

func TestDiscountModel_GetActual(t *testing.T) {
	discounts, err := getDiscountModel(t).GetActual()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(discounts)
}
