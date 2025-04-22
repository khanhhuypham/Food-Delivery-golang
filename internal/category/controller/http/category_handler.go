package category_http

import (
	category_dto "Food-Delivery/entity/dto/category"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryService interface {
	Create(ctx context.Context, cate *category_dto.CreateDto) error
	FindAll(ctx context.Context, paging *common.Paging, filter *category_dto.QueryDTO) ([]model.Category, error)
	FindOneById(ctx context.Context, id int) (*model.Category, error)
	Update(ctx context.Context, id int, dto *category_dto.CreateDto) error
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

		var dto category_dto.CreateDto
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		if err := handler.cateService.Create(ctx.Request.Context(), &dto); err != nil {
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
		var query category_dto.QueryDTO
		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		list, err := handler.cateService.FindAll(ctx.Request.Context(), &paging, &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(list, paging))
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

		var dto category_dto.CreateDto
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
