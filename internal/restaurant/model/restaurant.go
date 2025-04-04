package restaurant_model

import (
	"Food-Delivery/pkg/common"
	"errors"
	"github.com/google/uuid"
	"strings"
)

type QueryDTO struct {
	OwnerId *string `json:"ownerId"`
	CityId  *int    `json:"cityId"`
	Status  *string `json:"status"`
}

type Restaurant struct {
	common.SQLModel
	OwnerId uuid.UUID `json:"ownerId" gorm:"column:owner_id;"`
	Name    string    `json:"name" gorm:"column:name;"`
	Address string    `json:"address" gorm:"column:address;"`
	CityId  *int      `json:"cityId" gorm:"column:city_id;"`
	Lat     float64   `json:"lat" gorm:"column:lat;"`
	Lng     float64   `json:"lng" gorm:"column:lng;"`
	// Cover            json.RawMessage `json:"cover" gorm:"column:cover;"`
	// Logo             json.RawMessage `json:"logo" gorm:"column:logo;"`
	ShippingFeePerKm float64   `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km;"`
	Status           string    `json:"status" gorm:"column:status;"`
	CategoryId       uuid.UUID `json:"category_id" gorm:"column:category_id;type:char(36)"`
	//Category         *Category `json:"category" gorm:"foreignKey:CategoryId;references:Id;"`
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
