package httpApi

import (
	"CoffeLand/app/core/data"
	"CoffeLand/app/core/usecases"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type UserAPI struct {}

//products
func(api *UserAPI) GetProductsByFilters(apiCtx ApiContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		params := struct {
			ProductType string `json:"productType"`
			Volume      uint64 `json:"volume"`
			Name        string `json:"name"`
		}{}

		returnParams := struct {
			Products []interface{} `json:"products"`
			Error    error
		}{
			Products: []interface{}{},
			Error:    nil,
		}

		if ctx.Request.Method != http.MethodGet {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = errors.New("wrong method")
			return
		}

		err := json.NewDecoder(ctx.Request.Body).Decode(params)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		productUsecase := usecases.NewProductUsecases(apiCtx.Core)

		products, err := productUsecase.GetByFilters(data.ProductType(params.ProductType), params.Volume, params.Name)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Products = []interface{}{products}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

func(api *UserAPI) GetAllProducts(apiCtx ApiContext) func(*gin.Context) {
	return func(ctx *gin.Context) {

		returnParams := struct {
			Products []interface{} `json:"products"`
			Error    error
		}{
			Products: []interface{}{},
			Error:    nil,
		}

		if ctx.Request.Method != http.MethodGet {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = errors.New("wrong method")
			return
		}

		productUsecase := usecases.NewProductUsecases(apiCtx.Core)

		products, err := productUsecase.GetAll()
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Products = []interface{}{products}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

// discounts
func(api *UserAPI) GetDiscountByNameLike(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := struct {
			Name string `json:"name"`
		}{}

		returnParams := struct {
			Discounts []interface{} `json:"discounts"`
			Error    error
		}{
			Discounts: []interface{}{},
			Error:    nil,
		}

		err := json.NewDecoder(ctx.Request.Body).Decode(params)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		discountUsecase := usecases.NewDiscountUsecases(apiCtx.Core)

		discounts, err := discountUsecase.GetByNameLike(params.Name)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Discounts = []interface{}{discounts}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

func(api *UserAPI) GetActualDiscounts(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		returnParams := struct {
			Discounts []interface{} `json:"discounts"`
			Error    error
		}{
			Discounts: []interface{}{},
			Error:    nil,
		}

		if ctx.Request.Method != http.MethodGet {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = errors.New("wrong method")
			return
		}

		discountUsecase := usecases.NewDiscountUsecases(apiCtx.Core)

		discounts, err := discountUsecase.GetActual()
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Discounts = []interface{}{discounts}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

func(api *UserAPI) GetDiscountsByProductType(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := struct {
			ProductType string `json:"productType"`
		}{}

		returnParams := struct {
			Discounts []interface{} `json:"discounts"`
			Error    error
		}{
			Discounts: []interface{}{},
			Error:    nil,
		}

		err := json.NewDecoder(ctx.Request.Body).Decode(params)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		discountUsecase := usecases.NewDiscountUsecases(apiCtx.Core)

		discounts, err := discountUsecase.GetByProductType(data.ProductType(params.ProductType))
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Discounts = []interface{}{discounts}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

//blog posts
func(api *UserAPI) GetBlogPostByNameLike(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := struct {
			Name string `json:"name"`
		}{}

		returnParams := struct {
			Posts []interface{} `json:"posts"`
			Error    error
		}{
			Posts: []interface{}{},
			Error:    nil,
		}

		err := json.NewDecoder(ctx.Request.Body).Decode(params)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		postsUsecase := usecases.NewBlogPostUsecases(apiCtx.Core)

		posts, err := postsUsecase.GetByTitleLike(params.Name)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Posts = []interface{}{posts}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

func(api *UserAPI) GetAllBlogPosts(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		returnParams := struct {
			Posts []interface{} `json:"posts"`
			Error    error
		}{
			Posts: []interface{}{},
			Error:    nil,
		}

		if ctx.Request.Method != http.MethodGet {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = errors.New("wrong method")
			return
		}

		postsUsecase := usecases.NewBlogPostUsecases(apiCtx.Core)

		posts, err := postsUsecase.GetAll()
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Posts = []interface{}{posts}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

// place info

func(api *UserAPI) GetPlaceInfoByID(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := struct {
			ID string `json:"id"`
		}{}

		returnParams := struct {
			Info []interface{} `json:"info"`
			Error    error
		}{
			Info: []interface{}{},
			Error:    nil,
		}

		err := json.NewDecoder(ctx.Request.Body).Decode(params)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		infoUsecase := usecases.NewPlaceInfoUsecases(apiCtx.Core)

		info, err := infoUsecase.GetByID(params.ID)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Info = []interface{}{info}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

func(api *UserAPI) GetPlaceInfoByAddressLike(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := struct {
			Address string `json:"address"`
		}{}

		returnParams := struct {
			Info []interface{} `json:"info"`
			Error    error
		}{
			Info: []interface{}{},
			Error:    nil,
		}

		err := json.NewDecoder(ctx.Request.Body).Decode(params)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		infoUsecase := usecases.NewPlaceInfoUsecases(apiCtx.Core)

		info, err := infoUsecase.GetByAddressLike(params.Address)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.Info = []interface{}{info}

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

func(api *UserAPI) IsOpenPlaceInfo(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := struct {
			ID string `json:"id"`
		}{}

		returnParams := struct {
			IsOpen bool `json:"isOpen"`
			Error    error
		}{
			IsOpen: false,
			Error:    nil,
		}

		err := json.NewDecoder(ctx.Request.Body).Decode(params)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		infoUsecase := usecases.NewPlaceInfoUsecases(apiCtx.Core)

		isOpen, err := infoUsecase.IsOpen(params.ID)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		returnParams.IsOpen = isOpen

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}
