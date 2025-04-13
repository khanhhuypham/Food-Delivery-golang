package order_dto

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
	"errors"
)

type CreateDTO struct {
	UserId       int     `json:"user_id"`
	RestaurantId int     `json:"restaurant_id"`
	Description  *string `json:"description"`
}

func (dto *CreateDTO) Validate() error {

	if dto.UserId <= 0 {
		return common.ErrBadRequest(errors.New("user not found"))
	}

	if dto.RestaurantId <= 0 {
		return common.ErrBadRequest(errors.New("restaurant not found"))
	}

	return nil
}

type UpdateDTO struct {
	Status constant.OrderStatus `json:"status" form:"status"`
}

func (dto *UpdateDTO) Validate() error {

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
