package user_http

import (
	usermodel "Food-Delivery/internal/user/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService interface {
	Signup(ctx context.Context, dto *usermodel.UserCreate) error
	SignIn(ctx context.Context, dto *usermodel.UserLogin) (*utils.Token, error)
	DeleteUserById(ctx context.Context, id int) error
}

type userHandler struct {
	userService UserService
	hasher      *utils.Hasher
}

func NewUserHandler(userService UserService, hasher *utils.Hasher) *userHandler {
	return &userHandler{userService: userService, hasher: hasher}
}

func (handler *userHandler) Signup() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var dto usermodel.UserCreate

		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err).WithDebug(err.Error()))
		}

		if err := handler.userService.Signup(ctx, &dto); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.Response(dto.Id))
	}
}

func (handler *userHandler) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto usermodel.UserLogin

		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		token, err := handler.userService.SignIn(ctx, &dto)

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		ctx.JSON(http.StatusOK, common.Response(token))
	}
}
func (handler *userHandler) GetProfileAPI() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// lấy thông tin Requester
		requester := ctx.MustGet(common.KeyRequester).(common.Requester)
		ctx.JSON(http.StatusOK, common.Response(requester))

	}
}

func (handler *userHandler) DeleteUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := handler.hasher.Decode(ctx.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.userService.DeleteUserById(ctx, id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{"delete successfully": ctx.Param("id")})
	}
}
