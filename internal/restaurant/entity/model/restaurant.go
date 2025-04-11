package model

import (
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

type QueryDTO struct {
	OwnerId *string `json:"ownerId"`
	CityId  *int    `json:"cityId"`
	Status  *string `json:"status"`
}

type Restaurant struct {
	common.SQLModel
	OwnerId          int        `json:"ownerId" gorm:"column:owner_id;"`
	Name             string     `json:"name" gorm:"column:name;"`
	Address          string     `json:"address" gorm:"column:address;"`
	CityId           *int       `json:"cityId" gorm:"column:city_id;"`
	Lat              float64    `json:"lat" gorm:"column:lat;"`
	Lng              float64    `json:"lng" gorm:"column:lng;"`
	ShippingFeePerKm float64    `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km;"`
	Status           string     `json:"status" gorm:"column:status;"`
	MenuItems        []MenuItem `json:"menu_items" gorm:"foreignKey:RestaurantId;references:Id;"`
	Orders           []Order    `json:"orders" gorm:"foreignKey:RestaurantId;references:Id;"`
	//Category         *Category  `json:"category" gorm:"foreignKey:CategoryId;references:Id;"`
}

func (Restaurant) TableName() string {
	return "restaurant"
}

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusPending  = "pending"
	StatusDeleted  = "deleted"
)

func (r *Restaurant) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	r.Address = strings.TrimSpace(r.Address)

	if r.Name == "" {
		return common.ErrBadRequest(errors.New("restaurant name is required"))
	}

	if r.Address == "" {
		return common.ErrBadRequest(errors.New("restaurant address is required"))
	}

	if r.Status != StatusActive && r.Status != StatusInactive &&
		r.Status != StatusPending && r.Status != StatusDeleted {
		return common.ErrBadRequest(errors.New("status must be in (active, inactive, pending, deleted)"))

	}

	return nil
}
