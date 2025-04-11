package model

import "Food-Delivery/pkg/common"

const OrderItemEntity = "order item"

type OrderItem struct {
	common.SQLModel
	OrderId    int     `json:"order_id" gorm:"column:order_id;"`
	MenuItemId int     `json:"menu_item_id" gorm:"column:menu_item_id;"`
	Quantity   int     `json:"quantity" gorm:"column:quanity;"`
	Note       *string `json:"note" gorm:"column:note;"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
