package dto

import (
	restaurant_model "Food-Delivery/internal/restaurant/entity/model"
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

// DTO = Data Transfer Object
type RestaurantCreateDTO struct {
	Name   *string  `json:"name"`
	Addr   *string  `json:"addr"`
	CityId *int     `json:"cityId"`
	Lat    *float64 `json:"lat"`
	Lng    *float64 `json:"lng"`
	// Cover            *json.RawMessage `json:"cover"`
	// Logo             *json.RawMessage `json:"logo"`
	ShippingFeePerKm *float64 `json:"shippingFeePerKm"`
	Status           *string  `json:"status"`
}

func (r *RestaurantCreateDTO) Validate() error {
	if st := r.Status; st != nil &&
		*st != restaurant_model.StatusActive && *st != restaurant_model.StatusInactive &&
		*st != restaurant_model.StatusPending && *st != restaurant_model.StatusDeleted {
		return common.ErrBadRequest(errors.New("status must be in (active, inactive, pending, deleted)"))
	}

	if str := r.Name; str != nil {
		*r.Name = strings.TrimSpace(*str)

		if len(*r.Name) == 0 {
			return common.ErrBadRequest(errors.New("restaurant name is empty"))
		}
	}

	if addr := r.Addr; addr != nil {
		*r.Addr = strings.TrimSpace(*addr)

		if len(*r.Addr) == 0 {
			return common.ErrBadRequest(errors.New("restaurant address is required"))
		}
	}

	return nil
}
