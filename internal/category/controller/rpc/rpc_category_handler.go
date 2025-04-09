package rpc_category_handler

import (
	categorymodel "Food-Delivery/internal/category/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RPCCategoryService interface {
	FindAllByIds(ctx context.Context, ids []int) ([]categorymodel.Category, error)
}

type rpcHandler struct {
	rpcService RPCCategoryService
}

func NewRPCCategoryHandler(service RPCCategoryService) *rpcHandler {
	return &rpcHandler{service}
}

func (handler *rpcHandler) GetByIds() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto categorymodel.RPCRequestDTO

		if err := ctx.ShouldBindJSON(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		categories, err := handler.rpcService.FindAllByIds(ctx.Request.Context(), dto.Ids)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(categories))
	}
}
