package http

import (
	usermodel "Food-Delivery/internal/user/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService interface {
	Register(ctx context.Context, userCreate *usermodel.UserCreate) error
	Login(ctx context.Context, credentials *usermodel.UserLogin) (*utils.Token, error)
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
			panic(common.ErrBadRequest(err))
		}

		if err := handler.userService.Register(ctx, &dto); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"user": dto,
		})
	}
}

func (handler *userHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto usermodel.UserLogin
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		token, err := handler.userService.Login(ctx, &dto)

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
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
