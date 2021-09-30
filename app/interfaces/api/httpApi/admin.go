package httpApi

import (
	"CoffeLand/app/core/data"
	"CoffeLand/app/core/usecases"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminAPI struct { }

type ApiContext struct {
	Core data.Core
	ExecutorId string
}

type DiscountParams struct {
	ID string `json:"id"`
	Name string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	ProductType string `json:"productType"`
	Description string `json:"description"`
}

type ReturnParams struct {
	Id string `json:"id"`
	Error error `json:"error"`
}

func(api *AdminAPI) PutDiscount(apiCtx ApiContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		params := &DiscountParams{}
		returnParams := ReturnParams{
			Id:    "",
			Error: nil,
		}

		if ctx.Request.Method != http.MethodPost {
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

		adminUsecase, err := usecases.NewAdministratorUsecases(apiCtx.Core, apiCtx.ExecutorId)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		id, err := adminUsecase.PutDiscount(data.Discount{
			ID:          params.ID,
			Name:        params.Name,
			StartDate:   params.StartDate,
			EndDate:     params.EndDate,
			ProductType: data.ProductType(params.ProductType),
			Description: params.Description,
		})

		returnParams.Id = id

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

type ProductParams struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ProductType string `json:"productType"`
	Price uint64 `json:"price"`
	Volume uint64 `json:"volume"`
	VolumeUnits string `json:"volumeUnits"`
}

func(api *AdminAPI) PutProduct(apiCtx ApiContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		params := &ProductParams{}
		returnParams := ReturnParams{
			Id:    "",
			Error: nil,
		}

		if ctx.Request.Method != http.MethodPost {
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

		adminUsecase, err := usecases.NewAdministratorUsecases(apiCtx.Core, apiCtx.ExecutorId)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		id, err := adminUsecase.PutProduct(data.Product{
			ID:          params.ID,
			Name:        params.Name,
			ProductType: data.ProductType(params.ProductType),
			Price:       params.Price,
			Volume:      params.Volume,
			VolumeUnits: params.VolumeUnits,
		})

		returnParams.Id = id

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

type AdministratorParams struct {
	ID string `gorm:"primarykey"`
	FirstName string
	LastName string
	Patronymic string
	Password string
	Access string
}

func(api *AdminAPI) PutAdministrator(apiCtx ApiContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		params := &AdministratorParams{}
		returnParams := ReturnParams{
			Id:    "",
			Error: nil,
		}

		if ctx.Request.Method != http.MethodPost {
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

		adminUsecase, err := usecases.NewAdministratorUsecases(apiCtx.Core, apiCtx.ExecutorId)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		id, err := adminUsecase.PutAdministrator(data.Administrator{
			ID:         params.ID,
			FirstName:  params.FirstName,
			LastName:   params.LastName,
			Patronymic: params.Patronymic,
			Password:   params.Password,
			Access:     data.AdministratorAccess(params.Access),
			Hash:       nil,
		})

		returnParams.Id = id

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}

type BlogPostParams struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
	Date string `json:"date"`
}

func(api *AdminAPI) PutBlogPost(apiCtx ApiContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		params := &BlogPostParams{}
		returnParams := ReturnParams{
			Id:    "",
			Error: nil,
		}

		if ctx.Request.Method != http.MethodPost {
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

		adminUsecase, err := usecases.NewAdministratorUsecases(apiCtx.Core, apiCtx.ExecutorId)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			returnParams.Error = err
			return
		}

		id, err := adminUsecase.PutBlogPost(data.BlogPost{
			ID:    params.ID,
			Title: params.Title,
			Text:  params.Text,
			Date:  params.Date,
		})

		returnParams.Id = id

		ctx.Writer.WriteHeader(http.StatusOK)
		WriteResult(ctx.Writer, returnParams)
	}
}
