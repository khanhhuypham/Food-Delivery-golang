package menu_item_http_handler

import (
	menu_item_model "Food-Delivery/internal/menu_item/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MenuItemService interface {
	Create(ctx context.Context, menuItem *menu_item_model.MenuItemCreateDTO) (*menu_item_model.MenuItem, error)
	FindAll(ctx context.Context, paging *common.Paging, query *menu_item_model.QueryDTO) ([]menu_item_model.MenuItem, error)
	FindOneById(ctx context.Context, id int) (*menu_item_model.MenuItem, error)
	Update(ctx context.Context, id int, dto *menu_item_model.MenuItemCreateDTO) (*menu_item_model.MenuItem, error)
	Delete(ctx context.Context, id int) error
}

type menuItemHandler struct {
	menuItemService MenuItemService
}

func NewRestaurantHandler(menuItemService MenuItemService) *menuItemHandler {
	return &menuItemHandler{menuItemService}
}

func (handler *menuItemHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var item menu_item_model.MenuItemCreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&item); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		newItem, err := handler.menuItemService.Create(ctx.Request.Context(), &item)

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
		var query menu_item_model.QueryDTO
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

		var dto menu_item_model.MenuItemCreateDTO
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
