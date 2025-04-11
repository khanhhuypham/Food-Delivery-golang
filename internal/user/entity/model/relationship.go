package user_model

import "Food-Delivery/internal/order/entity/order_model"

type Order struct {
	Id           int                     `gorm:"column:id;"`
	UserId       int                     `json:"user_id" gorm:"column:user_id"`
	RestaurantId int                     `json:"restaurant_id" gorm:"column:restaurant_id;"`
	OrderItem    []OrderItem             `json:"orderItems" gorm:"foreignKey:OrderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalAmount  float64                 `json:"total_amount" gorm:"column:total_amount;"`
	Status       order_model.OrderStatus `json:"status" gorm:"column:status;"`
}

//==============================================================================

type OrderItem struct {
	Id       int     `gorm:"column:id;"`
	OrderId  int     `json:"order_id" gorm:"column:order_id;"`
	Quantity int     `json:"quantity" gorm:"column:quanity;"`
	Note     *string `json:"note" gorm:"column:note;"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
