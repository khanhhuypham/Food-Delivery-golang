package model

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

const RestaurantEntity = "restaurant"

type Restaurant struct {
	common.SQLModel
	OwnerId          int                       `json:"ownerId" gorm:"column:owner_id;"`
	Cover            *Media                    `json:"cover" gorm:"column:cover;"`
	Logo             *Media                    `json:"logo" gorm:"column:logo;"`
	Name             string                    `json:"name" gorm:"column:name;"`
	Email            *string                   `json:"email" gorm:"column:email;"`
	Phone            string                    `json:"phone" gorm:"column:phone;"`
	Address          string                    `json:"address" gorm:"column:address;"`
	CityId           *int                      `json:"cityId" gorm:"column:city_id;"`
	Lat              float64                   `json:"lat" gorm:"column:lat;"`
	Lng              float64                   `json:"lng" gorm:"column:lng;"`
	ShippingFeePerKm float64                   `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km;"`
	LikeCount        int                       `json:"like_count" gorm:"column:like_count; default:0"`
	Status           constant.RestaurantStatus `json:"status" gorm:"column:status;"`
	Description      *string                   `json:"description" gorm:"column:description;"`
	Items            []Item                    `json:"items" gorm:"foreignKey:RestaurantId;references:Id;"`
	Orders           []Order                   `json:"orders" gorm:"foreignKey:RestaurantId;references:Id;"`
	Rating           []Rating                  `json:"rating" gorm:"foreignKey:RestaurantId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Restaurant) TableName() string {
	return "restaurant"
}

func (r *Restaurant) Validate() error {
	r.Name = strings.TrimSpace(r.Name)

	r.Address = strings.TrimSpace(r.Address)

	if r.Name == "" {
		return common.ErrBadRequest(errors.New("restaurant name is required"))
	}

	if r.Address == "" {
		return common.ErrBadRequest(errors.New("restaurant address is required"))
	}

	if !r.Status.IsValid() {
		return common.ErrBadRequest(errors.New("invalid restaurant status"))
	}

	return nil
}
