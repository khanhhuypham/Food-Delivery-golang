package upload_http_handler

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/proto-buffer/gen/mediapb"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type MediaGRPC interface {
	UploadFiles(ctx context.Context, files []*mediapb.ImageUpload) ([]model.Media, error)
}

type mediaHandler struct {
	grpc MediaGRPC
}

func NewUploadHandler(grpc MediaGRPC) *mediaHandler {
	return &mediaHandler{grpc: grpc}
}
func (handler *mediaHandler) Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		const maxUploadSize = 500 * 1024 * 1024 // 500MB

		// Parse multipart form with a size limit (e.g., 32 MB)
		if err := ctx.Request.ParseMultipartForm(maxUploadSize); err != nil {
			ctx.JSON(400, common.ErrBadRequest(err))
			return
		}

		form := ctx.Request.MultipartForm
		files := form.File["files"] // Should be sent as `files[]` in form-data

		var images []*mediapb.ImageUpload

		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				ctx.JSON(400, common.ErrBadRequest(err))
				return
			}
			defer file.Close()

			dataBytes, err := io.ReadAll(file)

			if err != nil {
				ctx.JSON(400, common.ErrBadRequest(err))
				return
			}

			contentType := http.DetectContentType(dataBytes)

			image := &mediapb.ImageUpload{
				Filename:    fileHeader.Filename,
				Content:     dataBytes,
				ContentType: contentType,
			}

			images = append(images, image)
		}

		// (Optional) Log total files processed
		log.Printf("Total images uploaded: %d\n", len(images))

		// Now you can pass images to your gRPC handler if needed
		result, err := handler.grpc.UploadFiles(ctx, images)
		if err != nil {
			ctx.JSON(500, common.ErrInternal(err))
			return
		}

		ctx.JSON(200, common.Response(result)) // Return the processed image data
	}
}
