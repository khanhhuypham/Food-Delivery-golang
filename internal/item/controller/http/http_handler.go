package item_http_handler

import (
	item_dto "Food-Delivery/entity/dto/item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ItemService interface {
	Create(ctx context.Context, menuItem *item_dto.CreateDTO) (*model.Item, error)
	Update(ctx context.Context, id int, dto *item_dto.CreateDTO) (*model.Item, error)
	Delete(ctx context.Context, id int) error

	FindAll(ctx context.Context, paging *common.Paging, query *item_dto.QueryDTO) ([]item_dto.ItemDTO, error)
	FindOneById(ctx context.Context, id int) (*model.Item, error)
	FindTheMostPopularItem(ctx context.Context, paging *common.Paging) ([]item_dto.ItemDTO, error)
	FindTheMostRecommendedItem(ctx context.Context, paging *common.Paging) ([]item_dto.ItemDTO, error)
}

type itemHandler struct {
	itemService ItemService
}

func NewRestaurantHandler(menuItemService ItemService) *itemHandler {
	return &itemHandler{menuItemService}
}

func (handler *itemHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto item_dto.CreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		newItem, err := handler.itemService.Create(ctx.Request.Context(), &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(newItem))
	}
}

func (handler *itemHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto item_dto.CreateDTO
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		updatedItem, err := handler.itemService.Update(ctx.Request.Context(), id, &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(updatedItem))
	}
}

func (handler *itemHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.itemService.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (handler *itemHandler) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		//filter
		var query item_dto.QueryDTO
		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		items, err := handler.itemService.FindAll(ctx.Request.Context(), &paging, &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(items, paging))
	}
}

func (handler *itemHandler) FindOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		item, err := handler.itemService.FindOneById(ctx.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(item))
	}
}

func (handler *itemHandler) FindTheMostPopularItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()

		// check error from usecase layer
		items, err := handler.itemService.FindTheMostPopularItem(ctx.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(items, paging))
	}
}

func (handler *itemHandler) FindTheMostRecommendedItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()

		// check error from usecase layer
		items, err := handler.itemService.FindTheMostRecommendedItem(ctx.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(items, paging))
	}
}
