package order_model

import (
	"Food-Delivery/internal/order/model/relationship"
	usermodel "Food-Delivery/internal/user/model"
	"Food-Delivery/pkg/common"
)

const EntityName = "order"

type Order struct {
	common.SQLModel
	UserId       int                      `json:"user_id" gorm:"column:menu_id;"`
	User         usermodel.User           `json:"user" gorm:"foreignKey:UserId;references:Id;"`
	RestaurantId int                      `json:"restaurant_id" gorm:"column:restaurant_id;"`
	OrderItem    []relationship.OrderItem `json:"orderItems" gorm:"gorm:"foreignKey:OrderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalAmount  float64                  `json:"total_amount" gorm:"column:total_amount;"`
	Status       OrderStatus              `json:"status" gorm:"column:status;"`
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
