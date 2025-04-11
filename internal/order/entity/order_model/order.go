package order_model

import (
	"Food-Delivery/pkg/common"
)

const EntityName = "order"

type Order struct {
	common.SQLModel
	UserId       int         `json:"user_id" gorm:"column:user_id;"`
	RestaurantId int         `json:"restaurant_id" gorm:"column:restaurant_id;"`
	OrderItem    []OrderItem `json:"orderItems" gorm:"foreignKey:OrderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalAmount  float64     `json:"total_amount" gorm:"column:total_amount;"`
	Status       OrderStatus `json:"status" gorm:"column:status;"`
}

func (Order) TableName() string {
	return "order"
}

type OrderStatus string

const (
	Pending    OrderStatus = "Pending"
	InProgress OrderStatus = "In-Progress"
	Delivered  OrderStatus = "Delivered"
	Completed  OrderStatus = "Completed"
	Cancelled  OrderStatus = "Cancelled"
)
