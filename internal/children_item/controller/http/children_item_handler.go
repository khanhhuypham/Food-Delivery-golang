package children_item_http

import (
	children_item_dto "Food-Delivery/entity/dto/children_item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type childrenItemService interface {
	FindAll(ctx context.Context, filter *children_item_dto.QueryDTO) ([]model.ChildrenItem, error)
	FindOneById(ctx context.Context, id int) (*model.ChildrenItem, error)
	Create(ctx context.Context, dto *children_item_dto.CreateDTO) (*model.ChildrenItem, error)
	Update(ctx context.Context, id int, dto *children_item_dto.CreateDTO) (*model.ChildrenItem, error)
	Delete(ctx context.Context, id int) error
}

type childrenItemHandler struct {
	service childrenItemService
}

func NewCategoryHandler(service childrenItemService) *childrenItemHandler {
	return &childrenItemHandler{service}
}

func (handler *childrenItemHandler) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var query children_item_dto.QueryDTO
		if err := ctx.ShouldBind(&query); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		list, err := handler.service.FindAll(ctx.Request.Context(), &query)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(list))
	}
}

func (handler *childrenItemHandler) FindOneByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		//id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		cate, err := handler.service.FindOneById(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(cate))
	}
}

func (handler *childrenItemHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto children_item_dto.CreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		newData, err := handler.service.Create(ctx.Request.Context(), &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(newData))
	}
}

func (handler *childrenItemHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		//id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var dto children_item_dto.CreateDTO
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		updatedData, err := handler.service.Update(ctx.Request.Context(), id, &dto)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(updatedData))
	}
}

func (handler *childrenItemHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		//id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.service.Delete(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
