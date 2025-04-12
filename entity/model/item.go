package model

import "Food-Delivery/pkg/common"

const ItemEntity = "menu item"

type Item struct {
	common.SQLModel
	RestaurantId int     `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Name         string  `json:"name" gorm:"column:name;type:varchar(255);uniqueIndex:idx_item_name"`
	Description  *string `json:"description" gorm:"column:description;"`
	Price        float64 `json:"price" gorm:"column:price;"`
	Image        *common.Image
	OrderItems   []*OrderItem `json:"orderItems" gorm:"foreignKey:ItemId;references:Id"`
}

func (item *Item) TableName() string {
	return "item"
}
