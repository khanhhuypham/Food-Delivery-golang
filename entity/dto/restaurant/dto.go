package restaurant_dto

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

// DTO = Data Transfer Object
type CreateDTO struct {
	Name   *string  `json:"name"`
	Addr   *string  `json:"addr"`
	CityId *int     `json:"cityId"`
	Lat    *float64 `json:"lat"`
	Lng    *float64 `json:"lng"`
	// Cover            *json.RawMessage `json:"cover"`
	// Logo             *json.RawMessage `json:"logo"`
	ShippingFeePerKm *float64                   `json:"shippingFeePerKm"`
	Status           *constant.RestaurantStatus `json:"status"`
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

	if addr := dto.Addr; addr != nil {
		*dto.Addr = strings.TrimSpace(*addr)

		if len(*dto.Addr) == 0 {
			return common.ErrBadRequest(errors.New("restaurant address is required"))
		}
	}

	return nil
}
