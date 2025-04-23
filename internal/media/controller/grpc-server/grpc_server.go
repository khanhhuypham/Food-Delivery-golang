package media_grpc_server

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/proto-buffer/gen/mediapb"
	"context"
)

type MediaRepository interface {
	Create(ctx context.Context, media *model.Media) error
	FindDataWithCondition(ctx context.Context, condition map[string]any) (*model.Media, error)
	Delete(ctx context.Context, condition map[string]any) error
}

type mediaGrpcServer struct {
	mediapb.UnimplementedMediaServer
	repo MediaRepository
}

func NewCategoryGrpcServer(repo MediaRepository) *mediaGrpcServer {
	return &mediaGrpcServer{
		repo: repo,
	}
}

func (grpc *mediaGrpcServer) UploadImages(ctx context.Context, request *mediapb.UploadImagesRequest) (*mediapb.UploadImagesResponse, error) {
	//img, err := handler.s3Provider.UploadFile(ctx)
	//
	//if err != nil {
	//	panic(common.ErrBadRequest(err))
	//}
	//
	//if err := handler.mediaService.Create(ctx, img); err != nil {
	//	panic(err)
	//}

	return &mediapb.UploadImagesResponse{Data: nil}, nil
}
