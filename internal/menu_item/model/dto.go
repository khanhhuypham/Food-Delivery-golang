package menu_item_model

import (
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

type MenuItemCreateDTO struct {
	RestaurantId int           `json:"restaurant_id"`
	Name         string        `json:"name"`
	Image        *common.Image `json:"image"`
	Price        float32       `json:"price"`
	Description  *string       `json:"description"`
}

func (dto *MenuItemCreateDTO) Validate() error {

	dto.Name = strings.TrimSpace(dto.Name)

	if len(dto.Name) == 0 {
		return common.ErrBadRequest(errors.New("restaurant name is empty"))
	}

	if dto.Price <= 0 {
		return common.ErrBadRequest(errors.New("price must be greater than zero"))
	}

	return nil
}
