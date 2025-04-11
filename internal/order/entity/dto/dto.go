package dto

import (
	"Food-Delivery/internal/order/entity/order_model"
	"Food-Delivery/pkg/common"
	"errors"
)

type OrderCreateDTO struct {
	UserId       int     `json:"user_id"`
	RestaurantId int     `json:"restaurant_id"`
	Description  *string `json:"description"`
}

func (dto *OrderCreateDTO) Validate() error {

	if dto.UserId <= 0 {
		return common.ErrBadRequest(errors.New("user not found"))
	}

	if dto.RestaurantId <= 0 {
		return common.ErrBadRequest(errors.New("restaurant not found"))
	}

	return nil
}

type OrderUpdateDTO struct {
	Status order_model.OrderStatus `json:"status" form:"status"`
}

func (dto *OrderUpdateDTO) Validate() error {

	//dto.Name = strings.TrimSpace(dto.Name)
	//
	//if len(dto.Name) == 0 {
	//	return common.ErrBadRequest(errors.New("restaurant name is empty"))
	//}
	//
	//if dto.Price <= 0 {
	//	return common.ErrBadRequest(errors.New("price must be greater than zero"))
	//}

	return nil
}
