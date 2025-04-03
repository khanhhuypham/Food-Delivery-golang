package restaurant_http

import (
	restaurant_model "Food-Delivery/internal/restaurant/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type RestaurantService interface {
	Create(ctx context.Context, cate *restaurant_model.RestaurantCreateDTO) error
	FindAll(ctx context.Context, paging *common.Paging, filter *restaurant_model.QueryDTO) ([]restaurant_model.Restaurant, error)
	FindOneById(ctx context.Context, id uuid.UUID) (*restaurant_model.Restaurant, error)
	Update(ctx context.Context, id uuid.UUID, dto *restaurant_model.RestaurantCreateDTO) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type restaurantHandler struct {
	restaurantService RestaurantService
}

func NewRestaurantHandler(restaurantService RestaurantService) *restaurantHandler {
	return &restaurantHandler{restaurantService}
}

func (handler *restaurantHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var restaurant restaurant_model.RestaurantCreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&restaurant); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		if err := handler.restaurantService.Create(ctx.Request.Context(), &restaurant); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *restaurantHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		//filter
		var query restaurant_model.QueryDTO
		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		restaurants, err := handler.restaurantService.FindAll(ctx.Request.Context(), &paging, &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(restaurants, paging))
	}
}

func (handler *restaurantHandler) GetOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		cate, err := handler.restaurantService.FindOneById(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(cate))
	}
}

func (handler *restaurantHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto restaurant_model.RestaurantCreateDTO
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.restaurantService.Update(ctx.Request.Context(), id, &dto); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *restaurantHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.restaurantService.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
