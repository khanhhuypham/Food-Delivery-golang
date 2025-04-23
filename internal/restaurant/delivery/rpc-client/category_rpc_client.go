package category_rpc_client

import (
	"Food-Delivery/entity/model"
	"context"
	"fmt"

	"resty.dev/v3"
)

type CategoryRPCClient struct {
	catServiceURL string
}

func NewCategoryRPCClient(catServiceURL string) *CategoryRPCClient {
	return &CategoryRPCClient{catServiceURL: catServiceURL}
}

func (c *CategoryRPCClient) FindByIds(ctx context.Context, ids []int) ([]model.Category, error) {
	client := resty.New()

	type ResponseDTO struct {
		Data []model.Category `json:"data"`
	}

	var response ResponseDTO

	url := fmt.Sprintf("%s/find-by-ids", c.catServiceURL)

	_, err := client.R().
		SetBody(map[string]interface{}{
			"ids": ids,
		}).
		SetResult(&response).
		Post(url)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
