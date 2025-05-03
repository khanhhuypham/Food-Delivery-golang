package model

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
)

const OrderEntity = "menu item"

// ;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;
type Order struct {
	common.SQLModel
	UserId       int                  `json:"user_id" gorm:"column:user_id;not null"`
	RestaurantId int                  `json:"restaurant_id" gorm:"column:restaurant_id;not null"`
	DriverId     int                  `json:"driver_id" gorm:"column:driver_id;not null"`
	TotalAmount  float64              `json:"total_amount" gorm:"column:total_amount;not null"`
	Status       constant.OrderStatus `json:"status" gorm:"column:status;not null"`
	OrderItems   []*OrderItem         `json:"orderItems"  gorm:"foreignKey:OrderId;references:Id"`
}

/*
	1. Define the join table (OrderItem) explicitly.
	2 .Avoid using many2many: if you're using a custom join model like OrderItem.
 	3.Use foreignKey, joinForeignKey, references, and joinReferences as needed.
note:
	- You don’t use many2many: because OrderItem is a full join model (with additional fields like Quantity).
*/

func (order *Order) TableName() string {
	return "order"
}
