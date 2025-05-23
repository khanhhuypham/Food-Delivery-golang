package restaurant_http

import (
	restaurant_dto "Food-Delivery/entity/dto/restaurant"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RestaurantService interface {
	Create(ctx context.Context, cate *restaurant_dto.CreateDTO) error

	Update(ctx context.Context, id int, dto *restaurant_dto.CreateDTO) error
	Delete(ctx context.Context, id int) error

	FindAll(ctx context.Context, paging *common.Paging, filter *restaurant_dto.QueryDTO) ([]model.Restaurant, *restaurant_dto.Statistic, error)
	FindOneById(ctx context.Context, id int) (*model.Restaurant, error)
	FindTheMostPopularRestaurant(ctx context.Context, paging *common.Paging) ([]model.Restaurant, error)
	FindTheMostRecommendedRestaurant(ctx context.Context, paging *common.Paging) ([]model.Restaurant, error)
}

type CategoryGRPCClient interface {
	FindByIds(ctx context.Context, ids []int64) ([]model.Category, error)
}

type restaurantHandler struct {
	restaurantService  RestaurantService
	categoryGRPCClient CategoryGRPCClient
}

func NewRestaurantHandler(restaurantService RestaurantService, categoryGRPCClient CategoryGRPCClient) *restaurantHandler {
	return &restaurantHandler{
		restaurantService,
		categoryGRPCClient,
	}
}

func (handler *restaurantHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto restaurant_dto.CreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		if err := handler.restaurantService.Create(ctx.Request.Context(), &dto); err != nil {
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
		var query restaurant_dto.QueryDTO
		if err := ctx.ShouldBindQuery(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		restaurants, statistic, err := handler.restaurantService.FindAll(ctx.Request.Context(), &paging, &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPagingAndStatistic(
			restaurants,
			statistic,
			paging,
		))
	}
}

func (handler *restaurantHandler) GetOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		//categories, err := handler.categoryGRPCClient.FindByIds(ctx, []int64{1, 3})

		//for i, v := range categories {
		//	b, _ := json.MarshalIndent(v, "", "  ")
		//	fmt.Printf("%d: %s\n", i, b)
		//}

		if err != nil {
			// return restaurants, nil (if category is not important, we can return restaurants)
			panic(common.ErrInternal(err))
		}

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

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto restaurant_dto.CreateDTO
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

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.restaurantService.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (handler *restaurantHandler) FindTheMostPopularRestaurant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()

		// check error from usecase layer
		list, err := handler.restaurantService.FindTheMostPopularRestaurant(ctx.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		var data []restaurant_dto.RestaurantDTO
		for _, restaurant := range list {
			data = append(data, *restaurant.ToRestaurantDTO())
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(data, paging))
	}
}

func (handler *restaurantHandler) FindTheMostRecommendedRestaurant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()

		// check error from usecase layer
		list, err := handler.restaurantService.FindTheMostRecommendedRestaurant(ctx.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		var data []restaurant_dto.RestaurantDTO
		for _, restaurant := range list {
			data = append(data, *restaurant.ToRestaurantDTO())
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(data, paging))
	}
}
