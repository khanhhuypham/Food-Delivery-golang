package item_optional_http

import (
	item_optional_dto "Food-Delivery/entity/dto/item_optional"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type itemOptionalService interface {
	Create(ctx context.Context, dto *item_optional_dto.CreateDTO) (*model.Optional, error)
	Update(ctx context.Context, id int, dto *item_optional_dto.CreateDTO) (*model.Optional, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, restaurantId int) ([]model.Optional, error)
	FindOneById(ctx context.Context, id int) (*model.Optional, error)
}

type itemOptionalHandler struct {
	service itemOptionalService
}

func NewItemOptionalHandler(service itemOptionalService) *itemOptionalHandler {
	return &itemOptionalHandler{service}
}

func (handler *itemOptionalHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto item_optional_dto.CreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		newItem, err := handler.service.Create(ctx.Request.Context(), &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(newItem))
	}
}

func (handler *itemOptionalHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto item_optional_dto.CreateDTO
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		updatedItem, err := handler.service.Update(ctx.Request.Context(), id, &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(updatedItem))
	}
}

func (handler *itemOptionalHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.service.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (handler *itemOptionalHandler) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type queryDTO struct {
			RestaurantId int `form:"restaurant_id"`
		}

		var query queryDTO

		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		items, err := handler.service.FindAll(ctx.Request.Context(), query.RestaurantId)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(items))
	}
}

func (handler *itemOptionalHandler) FindOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		item, err := handler.service.FindOneById(ctx.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(item))
	}
}
