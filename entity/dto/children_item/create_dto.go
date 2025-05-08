package children_item_dto

import (
	"Food-Delivery/pkg/common"
)

type CreateDTO struct {
	RestaurantId int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	OptionalId   *int          `json:"-" gorm:"column:optional_id;"`
	Name         string        `json:"name" gorm:"column:name;"`
	Image        *common.Image `json:"image" gorm:"column:image;"`
	Price        float32       `json:"price" gorm:"column:price;"`
	Description  *string       `json:"description" gorm:"column:description;"`
}

func (dto *CreateDTO) Validate() error {

	return nil
}
