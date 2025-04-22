package category_grpc_client

import (
	"Food-Delivery/entity/model"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"resty.dev/v3"
)

type categoryGRPCClient struct {
	catGRPCServerURL string
	connect          *grpc.ClientConn
	client           category.CategoryClient
}

func NewCategoryGRPCClient(catGRPCServerURL string) *categoryGRPCClient {
	connect, err := grpc.NewClient(catGRPCServerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := category.NewCategoryClient(connect)

	return &categoryGRPCClient{
		catGRPCServerURL: catGRPCServerURL,
		connect:          connect,
		client:           client,
	}
}

func (grpc *categoryGRPCClient) FindByIds(ctx context.Context, ids []int) ([]model.Category, error) {
	//grpc.client.
	return nil, nil
}
