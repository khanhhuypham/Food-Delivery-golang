package rating_http_handler

import (
	rating_dto "Food-Delivery/entity/dto/rating"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RatingService interface {
	Update(ctx context.Context, dto *rating_dto.CreateDTO) error
}

type ratingHandler struct {
	ratingService RatingService
}

func NewRatingHandler(service RatingService) *ratingHandler {
	return &ratingHandler{service}
}

func (handler *ratingHandler) Like() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto rating_dto.CreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBindQuery(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		if err := handler.ratingService.Update(ctx.Request.Context(), &dto); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *ratingHandler) Comment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dto rating_dto.CreateDTO
		//error occurs from binding json data into struct data
		if err := ctx.ShouldBind(&dto); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// check error from usecase layer
		if err := handler.ratingService.Update(ctx.Request.Context(), &dto); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}

func (handler *ratingHandler) SetScore() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, common.Response("ok"))
	}
}
