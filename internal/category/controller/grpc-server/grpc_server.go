package category_grpc

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/proto-buffer/proto/category"
	"context"
)

type CategoryRepository interface {
	FindAllByIds(ctx context.Context, ids []int, keys ...string) ([]model.Category, error)
}

type CategoryGrpcServer interface {
	category.UnimplementedCategoryServiceServer
	repo CategoryRepository
}

func NewCategoryGrpcServer(repo CategoryRepository) *CategoryGrpcServer {
	return &CategoryGrpcServer{
		repo: repo,
	}
}

func (grpc *CategoryGrpcServer) GetCategoriesByIds(ctx context.Context, req *category.GetCatIdsRequest) (*category.CatIdsResp, error) {
	cat,err := grpc.repo.FindAllByIds(ctx,req.ids)



	return nil, nil
}
