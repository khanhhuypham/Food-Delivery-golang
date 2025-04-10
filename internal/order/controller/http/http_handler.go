package order_http_handler

import (
	order_model "Food-Delivery/internal/order/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderService interface {
	Create(ctx context.Context, data *order_model.OrderCreateDTO) error
	FindAll(ctx context.Context, paging *common.Paging, query *order_model.QueryDTO) ([]order_model.Order, error)
	FindOneById(ctx context.Context, id int) (*order_model.Order, error)
	ChangeStatus(ctx context.Context, id int, dto *order_model.OrderUpdateDTO) (*order_model.Order, error)
}

type orderHandler struct {
	orderService OrderService
}

func NewOrderHandler(orderService OrderService) *orderHandler {
	return &orderHandler{orderService}
}

func (handler *orderHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto order_model.OrderCreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		if err := handler.orderService.Create(ctx.Request.Context(), &dto); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *orderHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		//filter
		var query order_model.QueryDTO
		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		restaurants, err := handler.orderService.FindAll(ctx.Request.Context(), &paging, &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(restaurants, paging))
	}
}

func (handler *orderHandler) GetOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		cate, err := handler.orderService.FindOneById(ctx.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(cate))
	}
}

func (handler *orderHandler) ChangeStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto order_model.OrderUpdateDTO

		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		data, err := handler.orderService.ChangeStatus(ctx.Request.Context(), id, &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(data))
	}
}
