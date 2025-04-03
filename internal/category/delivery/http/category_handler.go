package category_http

import (
	categorymodel "Food-Delivery/internal/category/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryService interface {
	Create(ctx context.Context, cate *categorymodel.CategoryCreateDto) error
	FindAll(ctx context.Context, paging *common.Paging, filter *categorymodel.QueryDTO) ([]categorymodel.Category, error)
	FindOneById(ctx context.Context, id int) (*categorymodel.Category, error)
	Update(ctx context.Context, id int, dto *categorymodel.CategoryCreateDto) error
	Delete(ctx context.Context, id int) error
}

type categoryHandler struct {
	cateService CategoryService
}

func NewCategoryHandler(cateService CategoryService) *categoryHandler {
	return &categoryHandler{cateService}
}

func (handler *categoryHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var cate categorymodel.CategoryCreateDto
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&cate); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		if err := handler.cateService.Create(ctx.Request.Context(), &cate); err != nil {
			panic(err)
		}

		////Encode id trước trả ra cho client
		//cate.FakeId = handler.hasher.Encode(cate.Id, common.DBTypePlace)

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *categoryHandler) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		//filter
		var query categorymodel.QueryDTO
		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		places, err := handler.cateService.FindAll(ctx.Request.Context(), &paging, &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(places, paging))
	}
}

func (handler *categoryHandler) FindOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		//id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		cate, err := handler.cateService.FindOneById(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(cate))
	}
}

func (handler *categoryHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		//id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto categorymodel.CategoryCreateDto
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.cateService.Update(ctx.Request.Context(), id, &dto); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *categoryHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		//id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.cateService.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
