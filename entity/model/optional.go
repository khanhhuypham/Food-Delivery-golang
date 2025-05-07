package model

import "Food-Delivery/pkg/common"

type Optional struct {
	common.SQLModel
	RestaurantId  int            `json:"restaurant_id" gorm:"column:restaurant_id;not null"`
	Name          string         `json:"name" gorm:"column:name;not null"`
	Description   *string        `json:"description" gorm:"column:description;"`
	ItemId        int            `json:"item_id" gorm:"column:item_id;not null"`
	ChildrenItems []ChildrenItem `json:"children_items" gorm:"foreignKey:OptionalId;references:Id"`
}

func (Optional) TableName() string {
	return "optional"
}
