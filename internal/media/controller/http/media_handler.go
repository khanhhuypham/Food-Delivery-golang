package media_http

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/upload"
	"context"
	"github.com/gin-gonic/gin"
)

type MediaService interface {
	Create(ctx context.Context, media *model.Media) error
}

type mediaHandler struct {
	s3Provider   upload.UploadProvider
	mediaService MediaService
}

func NewMediaHandler(s3Provider upload.UploadProvider, mediaService MediaService) *mediaHandler {
	return &mediaHandler{
		s3Provider:   s3Provider,
		mediaService: mediaService,
	}
}

func (handler *mediaHandler) Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		img, err := handler.s3Provider.UploadFile(ctx)

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := handler.mediaService.Create(ctx, img); err != nil {
			panic(err)
		}

		ctx.JSON(200, common.Response(img))
	}
}
