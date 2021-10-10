package data

type ProductType string

const (
	ProductTypeCoffee = ProductType("coffee")
)

func(pt ProductType) IsValid() bool {
	switch pt {
	case ProductTypeCoffee:
		return true
	default:
		return false
	}
}

type Core struct {
	AdminRepo AdministratorRepository
	DiscountRepo DiscountsRepository
	ProductRepo ProductsRepository
	BlogPostRepo BlogPostRepository
	PlaceInfoRepo PlaceInfoRepository
}

