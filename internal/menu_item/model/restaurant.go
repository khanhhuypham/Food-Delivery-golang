package menu_item_model

import "Food-Delivery/pkg/common"

type Restaurant struct {
	common.SQLModel
	Name             string  `json:"name" gorm:"column:name;"`
	ShippingFeePerKm float64 `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km;"`
	Status           string  `json:"status" gorm:"column:status;"`
	CategoryId       int     `json:"category_id" gorm:"column:category_id"`
}

func (Restaurant) TableName() string {
	return "restaurant"
}
