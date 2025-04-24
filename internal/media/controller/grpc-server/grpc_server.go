package media_grpc_server

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/upload"
	"Food-Delivery/proto-buffer/gen/mediapb"
	"context"
)

type MediaService interface {
	Create(ctx context.Context, media *model.Media) error
}

type mediaGrpcServer struct {
	mediapb.UnimplementedMediaServer
	s3Provider   upload.UploadProvider
	mediaService MediaService
}

func NewMediaGrpcServer(s3Provider upload.UploadProvider, mediaService MediaService) *mediaGrpcServer {
	return &mediaGrpcServer{
		s3Provider:   s3Provider,
		mediaService: mediaService,
	}
}

func (grpc *mediaGrpcServer) UploadImages(ctx context.Context, request *mediapb.UploadImagesRequest) (*mediapb.UploadImagesResponse, error) {

	fileArray, err := grpc.s3Provider.UploadFiles(ctx, request.Images)
	result := make([]*mediapb.MediaMessage, len(fileArray))
	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	for i, file := range fileArray {

		if err := grpc.mediaService.Create(ctx, file); err != nil {
			panic(err)
		}

		result[i] = &mediapb.MediaMessage{
			Id:     int64(file.Id),
			Url:    file.Url,
			Size:   file.Size,
			Width:  intToInt64Ptr(file.Width),
			Height: intToInt64Ptr(file.Height),
		}
	}

	return &mediapb.UploadImagesResponse{Data: result}, nil
}

func intToInt64Ptr(v *int) *int64 {
	if v == nil {
		return nil
	}
	val := int64(*v)
	return &val
}
