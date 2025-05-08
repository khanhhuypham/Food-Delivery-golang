package item_optional_dto

import "Food-Delivery/entity/model"

type ItemOptionalDTO struct {
	ID           int     `json:"id"`
	RestaurantId int     `json:"restaurant_id" `
	Name         string  `json:"name" `
	Description  *string `json:"description"`
	ItemId       int     `json:"item_id" `
}

type ItemOptionalDetailDTO struct {
	ID            int                  `json:"id"`
	RestaurantId  int                  `json:"restaurant_id" `
	Name          string               `json:"name" `
	Description   *string              `json:"description"`
	ItemId        int                  `json:"item_id" `
	ChildrenItems []model.ChildrenItem `json:"children_items" `
}
