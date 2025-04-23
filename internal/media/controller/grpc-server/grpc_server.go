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

	img, err := grpc.s3Provider.UploadFiles(ctx, request.Images)

	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	if err := grpc.mediaService.Create(ctx, img); err != nil {
		panic(err)
	}

	return &mediapb.UploadImagesResponse{Data: nil}, nil
}
