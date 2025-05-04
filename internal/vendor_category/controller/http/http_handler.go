package vendor_category_http_handler

import (
	vendor_category_dto "Food-Delivery/entity/dto/vendor_category"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
)

type VendorCategoryService interface {
	FindAll(ctx context.Context, restaurantId int) ([]model.VendorCategory, error)
	FindOneById(ctx context.Context, id int) (*model.VendorCategory, error)
	Create(ctx context.Context, dto *vendor_category_dto.CreateDTO) (*model.VendorCategory, error)
	Update(ctx context.Context, id int, dto *vendor_category_dto.UpdateDTO) (*model.VendorCategory, error)
	Delete(ctx context.Context, id int) error
}

type vendorCategoryHandler struct {
	vendorCategoryService VendorCategoryService
}

func NewVendorCategoryHandler(vendorCategoryService VendorCategoryService) *vendorCategoryHandler {
	return &vendorCategoryHandler{vendorCategoryService}
}

func (handler *vendorCategoryHandler) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		type queryDTO struct {
			RestaurantId int `form:"restaurant_id"`
		}

		var query queryDTO

		if err := ctx.ShouldBindQuery(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}
		// check error from usecase layer
		vendorCategories, err := handler.vendorCategoryService.FindAll(ctx.Request.Context(), query.RestaurantId)
		if err != nil {
			panic(err)
		}

		// Convert list of VendorCategory -> list of VendorCategoryDTO
		var list []vendor_category_dto.VendorCategoryDTO

		for _, v := range vendorCategories {
			var dto vendor_category_dto.VendorCategoryDTO
			copier.Copy(&dto, &v)
			dto.TotalItems = len(v.Items)
			dto.Items = nil
			list = append(list, dto)
		}

		ctx.JSON(http.StatusOK, common.Response(list))
	}
}

func (handler *vendorCategoryHandler) FindOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		cate, err := handler.vendorCategoryService.FindOneById(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(cate.ToVendorCategoryDetailDTO()))
	}
}

func (handler *vendorCategoryHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto vendor_category_dto.CreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		data, err := handler.vendorCategoryService.Create(ctx.Request.Context(), &dto)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(data))
	}
}

func (handler *vendorCategoryHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto vendor_category_dto.UpdateDTO
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		data, err := handler.vendorCategoryService.Update(ctx.Request.Context(), id, &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(data))
	}
}

func (handler *vendorCategoryHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.vendorCategoryService.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
