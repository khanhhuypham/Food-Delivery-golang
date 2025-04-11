package dto

import (
	"Food-Delivery/pkg/common"
	"errors"
)

// DTO = Data Transfer Object
type OrderItemCreateDTO struct {
	OrderId    int     `json:"order_id"`
	MenuItemId int     `json:"menu_item_id"`
	Quantity   int     `json:"quantity"`
	Note       *string `json:"note"`
}

func (dto *OrderItemCreateDTO) Validate() error {

	if dto.Quantity < 0 {
		return common.ErrBadRequest(errors.New("quantity couldn't be negative"))
	}

	if dto.OrderId <= 0 {
		return common.ErrBadRequest(errors.New("order not found"))
	}

	if dto.MenuItemId <= 0 {
		return common.ErrBadRequest(errors.New("menu item not found"))
	}

	return nil
}
