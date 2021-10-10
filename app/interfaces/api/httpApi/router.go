package httpApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(ctx ApiContext) {
	router := Router{Engine: gin.New()}
	router.Engine.Use(gin.Logger())

	router.Engine.GET("/", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})
	router.Engine.GET("/favicon.ico", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})

	bindRoutes(router.Engine, ctx)
}

func bindRoutes(router *gin.Engine, ctx ApiContext) {
	router.POST("/admin/put_product", AdminAPI{}.PutProduct(ctx))
	router.POST("/admin/put_discount", AdminAPI{}.PutDiscount(ctx))
	router.POST("/admin/put_administrator", AdminAPI{}.PutAdministrator(ctx))
	router.POST("/admin/put_place_info", AdminAPI{}.PutPlaceInfo(ctx))

	router.GET("/get_products_by_filters", UserAPI{}.GetProductsByFilters(ctx))
	router.GET("/get_all_products", UserAPI{}.GetAllProducts(ctx))
	router.GET("/get_discount_by_name_like", UserAPI{}.GetDiscountByNameLike(ctx))
	router.GET("/actual_discounts", UserAPI{}.GetActualDiscounts(ctx))
	router.GET("/discounts_by_product_type", UserAPI{}.GetDiscountsByProductType(ctx))
	router.GET("/place_info_by_id", UserAPI{}.GetPlaceInfoByID(ctx))
	router.GET("/place_info_by_address", UserAPI{}.GetPlaceInfoByAddressLike(ctx))
	router.GET("/place_info_is_open", UserAPI{}.IsOpenPlaceInfo(ctx))

	router.GET("/admin/get_products_by_filters", UserAPI{}.GetProductsByFilters(ctx))
	router.GET("/admin/get_all_products", UserAPI{}.GetAllProducts(ctx))
	router.GET("/admin/get_discount_by_name_like", UserAPI{}.GetDiscountByNameLike(ctx))
	router.GET("/admin/actual_discounts", UserAPI{}.GetActualDiscounts(ctx))
	router.GET("/admin/discounts_by_product_type", UserAPI{}.GetDiscountsByProductType(ctx))
}
