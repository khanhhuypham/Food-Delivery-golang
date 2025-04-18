package restaurant_dto

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

// DTO = Data Transfer Object
type CreateDTO struct {
	Name    *string      `json:"name"`
	Email   *string      `json:"email"`
	Phone   string       `json:"phone"`
	Address *string      `json:"address" gorm:"column:address;"`
	Cover   *model.Media `json:"cover" gorm:"column:cover;"`
	Logo    *model.Media `json:"logo" gorm:"column:image;"`
	//ShippingFeePerKm *float64                   `json:"shippingFeePerKm"`
	Description *string                    `json:"description" gorm:"column:description;"`
	Status      *constant.RestaurantStatus `json:"status"`
}

func (dto *CreateDTO) Validate() error {
	if status := dto.Status; status != nil && !dto.Status.IsValid() {
		return common.ErrBadRequest(errors.New("status invalid"))
	}

	if str := dto.Name; str != nil {
		*dto.Name = strings.TrimSpace(*str)

		if len(*dto.Name) == 0 {
			return common.ErrBadRequest(errors.New("restaurant name is empty"))
		}
	}

	if addr := dto.Address; addr != nil {
		*dto.Address = strings.TrimSpace(*addr)

		if len(*dto.Address) == 0 {
			return common.ErrBadRequest(errors.New("restaurant address is required"))
		}
	}

	return nil
}
