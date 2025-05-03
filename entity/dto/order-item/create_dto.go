package order_item_dto

import (
	"Food-Delivery/pkg/common"
	"errors"
)

// DTO = Data Transfer Object
type CreateDTO struct {
	OrderId  int     `json:"order_id" gorm:"column:order_id"`
	ItemId   int     `json:"item_id" gorm:"column:item_id"`
	Quantity int     `json:"quantity" gorm:"column:quantity"`
	Note     *string `json:"note" gorm:"column:note"`
}

func (CreateDTO) TableName() string {
	return "order_item"
}

func (dto *CreateDTO) Validate() error {

	if dto.Quantity < 0 {
		return common.ErrBadRequest(errors.New("quantity couldn't be negative"))
	}

	if dto.OrderId <= 0 {
		return common.ErrBadRequest(errors.New("order not found"))
	}

	if dto.ItemId <= 0 {
		return common.ErrBadRequest(errors.New("menu item not found"))
	}

	return nil
}
