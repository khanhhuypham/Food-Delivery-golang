package relationship

import (
	order_model "Food-Delivery/internal/order/model"
	order_item_model "Food-Delivery/internal/order_item/model"
	usermodel "Food-Delivery/internal/user/model"
)

type Order struct {
	Id           int                          `gorm:"column:id;"`
	UserId       int                          `gorm:"column:menu_id;"`
	User         usermodel.User               `gorm:"foreignKey:UserId;references:Id;"`
	RestaurantId int                          `gorm:"column:restaurant_id;"`
	OrderItem    []order_item_model.OrderItem `gorm:"gorm:"foreignKey:OrderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalAmount  float64                      `gorm:"column:total_amount;"`
	Status       order_model.OrderStatus      `gorm:"column:status;"`
}
