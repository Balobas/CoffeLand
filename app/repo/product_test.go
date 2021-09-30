package repo

import (
	"CoffeLand/app/core/data"
	"CoffeLand/app/interfaces/database"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func getProductModel(t *testing.T) data.ProductsRepository{
	db, err := database.MySQLDB()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return NewProductRepo(db)
}

func TestProductModel_GetByID(t *testing.T) {
	pm := getProductModel(t)

	product, err := pm.GetByID("dsss")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println()
	t.Log(product)


	_, err = pm.GetByID("adsadasd")
	if err != gorm.ErrRecordNotFound {
		t.Error("expected record not found error")
		t.FailNow()
	}
}

func TestProductModel_GetByType(t *testing.T) {
	pm := getProductModel(t)

	products, err := pm.GetByType("coffee")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println()
	t.Log(products)


	_, err = pm.GetByType("adsadasd")
	if err != gorm.ErrRecordNotFound {
		t.Error("expected record not found error")
		t.FailNow()
	}
}

func TestProductModel_GetByTypeAndVolume(t *testing.T) {
	pm := getProductModel(t)

	product, err := pm.GetByTypeAndVolume("coffee", 300)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println()
	t.Log(product)


	_, err = pm.GetByTypeAndVolume("coffee", 100)
	if err != gorm.ErrRecordNotFound {
		t.Error("expected record not found error")
		t.FailNow()
	}
}

func TestProductModel_GetByNameLike(t *testing.T) {
	pm := getProductModel(t)

	product, err := pm.GetByNameLike("k")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println()
	t.Log(product)


	_, err = pm.GetByNameLike("adsadasd")
	if err != gorm.ErrRecordNotFound {
		t.Error("expected record not found error")
		t.FailNow()
	}
}

func TestProductModel_Store(t *testing.T) {
	product := data.Product{
		ID:          "ooaooaoaoa",
		Name:        "cofein",
		ProductType:        "coffee",
		Price:       40,
		Volume:      150,
		VolumeUnits: "мл",
	}

	pm := getProductModel(t)

	if err := pm.Store(product); err != nil {
		t.Error(err)
		t.FailNow()
	}

}
