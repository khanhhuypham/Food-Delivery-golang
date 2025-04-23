package category_grpc_client

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/proto-buffer/gen/categorypb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type categoryGRPCClient struct {
	catGRPCServerURL string
	connect          *grpc.ClientConn
	client           categorypb.CategoryClient
}

func NewCategoryGRPCClient(catGRPCServerURL string) *categoryGRPCClient {
	connect, err := grpc.NewClient(
		catGRPCServerURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	client := categorypb.NewCategoryClient(connect)

	return &categoryGRPCClient{
		catGRPCServerURL: catGRPCServerURL,
		connect:          connect,
		client:           client,
	}
}

func (grpc *categoryGRPCClient) FindByIds(ctx context.Context, ids []int64) ([]model.Category, error) {

	resp, err := grpc.client.GetCategoriesByIds(ctx, &categorypb.GetCategoriesRequest{Ids: ids})

	if err != nil {
		return nil, err
	}

	result := make([]model.Category, len(resp.Data))

	for i, cat := range resp.Data {

		result[i] = model.Category{
			SQLModel: common.SQLModel{
				Id: int(cat.Id),
			},
			Name:   cat.Name,
			Status: constant.CategoryStatus(cat.Status),
		}
	}

	return result, nil
}
