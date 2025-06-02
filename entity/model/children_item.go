package model

import "Food-Delivery/pkg/common"

type ChildrenItem struct {
	common.SQLModel
	Image        *common.Image `json:"image" gorm:"column:image;"`
	RestaurantId int           `json:"restaurant_id" gorm:"column:restaurant_id;not null"`
	OptionalId   *int          `json:"optional_id" gorm:"column:optional_id"`
	Name         string        `json:"name" gorm:"column:name;not null"`
	Description  *string       `json:"description" gorm:"column:description;"`
	Active       bool          `json:"active" gorm:"column:active;default:true"`
	Price        float32       `json:"price" gorm:"column:price;not null"`
	OutOfStock   bool          `json:"out_of_stock" gorm:"column:out_of_stock;default:false"`
}

func (ChildrenItem) TableName() string {
	return "children_item"
}
