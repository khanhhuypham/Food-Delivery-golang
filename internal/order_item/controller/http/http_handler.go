package order_item_http_handler

import (
	order_item_model "Food-Delivery/internal/order_item/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderItemService interface {
	Create(ctx context.Context, dto *order_item_model.OrderItemCreateDTO) error
	FindAll(ctx context.Context, paging *common.Paging, query *order_item_model.QueryDTO) ([]order_item_model.OrderItem, error)
	FindOneById(ctx context.Context, id int) (*order_item_model.OrderItem, error)
	Update(ctx context.Context, id int, dto *order_item_model.OrderItemCreateDTO) error
	Delete(ctx context.Context, id int) error
}

type orderItemHandler struct {
	orderItemService OrderItemService
}

func NewOrderItemHandler(service OrderItemService) *orderItemHandler {
	return &orderItemHandler{service}
}

func (handler *orderItemHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto order_item_model.OrderItemCreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		if err := handler.orderItemService.Create(ctx.Request.Context(), &dto); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *orderItemHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		//filter
		var query order_item_model.QueryDTO
		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		restaurants, err := handler.orderItemService.FindAll(ctx.Request.Context(), &paging, &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(restaurants, paging))
	}
}

func (handler *orderItemHandler) GetOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		cate, err := handler.orderItemService.FindOneById(ctx.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(cate))
	}
}

func (handler *orderItemHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto order_item_model.OrderItemCreateDTO
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.orderItemService.Update(ctx.Request.Context(), id, &dto); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *orderItemHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.orderItemService.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
