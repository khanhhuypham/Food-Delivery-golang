package category_grpc_hanlder

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/proto-buffer/gen/categorypb"
	"context"
	"log"
)

type CategoryRepository interface {
	FindAllByIds(ctx context.Context, ids []int, keys ...string) ([]model.Category, error)
}

type categoryGrpcServer struct {
	categorypb.UnimplementedCategoryServer
	repo CategoryRepository
}

func NewCategoryGrpcServer(repo CategoryRepository) *categoryGrpcServer {
	return &categoryGrpcServer{
		repo: repo,
	}
}

func (grpc *categoryGrpcServer) GetCategoriesByIds(ctx context.Context, request *categorypb.GetCategoriesRequest) (*categorypb.GetCategoriesResponse, error) {
	ids := make([]int, len(request.Ids))
	for i, v := range request.Ids {
		ids[i] = int(v)
	}
	list, err := grpc.repo.FindAllByIds(ctx, ids)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]*categorypb.CategoryMessage, len(list))

	for i, item := range list {
		result[i] = &categorypb.CategoryMessage{
			Id:     int64(item.Id),
			Name:   item.Name,
			Status: string(item.Status),
		}
	}

	return &categorypb.GetCategoriesResponse{Data: result}, nil
}
