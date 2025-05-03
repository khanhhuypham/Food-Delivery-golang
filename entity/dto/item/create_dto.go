package item_dto

import (
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

type CreateDTO struct {
	RestaurantId     int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId       int           `json:"category_id" gorm:"column:category_id;"`
	VendorCategoryId int           `json:"vendor_category_id" gorm:"column:vendor_category_id;"`
	Name             string        `json:"name" gorm:"column:name;"`
	Image            *common.Image `json:"image" gorm:"column:image;"`
	Price            float32       `json:"price" gorm:"column:price;"`
	Description      *string       `json:"description" gorm:"column:description;"`
}

func (dto *CreateDTO) Validate() error {

	dto.Name = strings.TrimSpace(dto.Name)

	if len(dto.Name) == 0 {
		return common.ErrBadRequest(errors.New("restaurant name is empty"))
	}

	if dto.Price <= 0 {
		return common.ErrBadRequest(errors.New("price must be greater than zero"))
	}

	return nil
}
