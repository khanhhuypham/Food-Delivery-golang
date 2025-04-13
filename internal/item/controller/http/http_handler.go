package item_http_handler

import (
	menu_item_dto "Food-Delivery/entity/dto/item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ItemService interface {
	Create(ctx context.Context, menuItem *menu_item_dto.CreateDTO) (*model.Item, error)
	FindAll(ctx context.Context, paging *common.Paging, query *menu_item_dto.QueryDTO) ([]model.Item, error)
	FindOneById(ctx context.Context, id int) (*model.Item, error)
	Update(ctx context.Context, id int, dto *menu_item_dto.CreateDTO) (*model.Item, error)
	Delete(ctx context.Context, id int) error
}

type menuItemHandler struct {
	menuItemService ItemService
}

func NewRestaurantHandler(menuItemService ItemService) *menuItemHandler {
	return &menuItemHandler{menuItemService}
}

func (handler *menuItemHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto menu_item_dto.CreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		newItem, err := handler.menuItemService.Create(ctx.Request.Context(), &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(newItem))
	}
}

func (handler *menuItemHandler) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		//filter
		var query menu_item_dto.QueryDTO
		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		items, err := handler.menuItemService.FindAll(ctx.Request.Context(), &paging, &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(items, paging))
	}
}

func (handler *menuItemHandler) FindOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		item, err := handler.menuItemService.FindOneById(ctx.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(item))
	}
}

func (handler *menuItemHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto menu_item_dto.CreateDTO
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		updatedItem, err := handler.menuItemService.Update(ctx.Request.Context(), id, &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(updatedItem))
	}
}

func (handler *menuItemHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.menuItemService.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
