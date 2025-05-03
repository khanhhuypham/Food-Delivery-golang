package model

import "Food-Delivery/pkg/common"

const OrderItemEntity = "order item"

type OrderItem struct {
	common.SQLModel
	OrderId  int     `json:"order_id" gorm:"column:order_id;not null"`
	ItemId   int     `json:"item_id" gorm:"column:item_id;not null"`
	Order    *Order  `gorm:"foreignKey:OrderId;references:Id"`
	Item     *Item   `gorm:"foreignKey:ItemId;references:Id"`
	Quantity int     `json:"quantity" gorm:"column:quantity;not null"`
	Note     *string `json:"note" gorm:"column:note"`
}

func (orderItem OrderItem) TableName() string {
	return "order_item"
}
