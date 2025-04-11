package user_rpc

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService interface {
	FindById(ctx context.Context, id int) (*model.User, error)
}

type rpcUserHandler struct {
	userService UserService
	hasher      *utils.Hasher
}

func NewRPCUserHandler(userService UserService, hasher *utils.Hasher) *rpcUserHandler {
	return &rpcUserHandler{userService: userService, hasher: hasher}
}

func (handler *rpcUserHandler) IntrospectTokenRPC() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var bodyData struct {
			Token string `json:"token"`
		}

		if err := ctx.ShouldBindJSON(&bodyData); err != nil {
			panic(common.ErrBadRequest(err).WithDebug(err.Error()))
		}

		user, err := handler.userService.FindById(ctx, 5)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(user))
	}

}
