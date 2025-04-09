package rpc_category_client

import (
	restaurant_model "Food-Delivery/internal/restaurant/model"
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

func (c *CategoryRPCClient) FindByIds(ctx context.Context, ids []int) ([]restaurant_model.Category, error) {
	client := resty.New()

	type ResponseDTO struct {
		Data []restaurant_model.Category `json:"data"`
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
