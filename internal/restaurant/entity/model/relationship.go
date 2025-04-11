package model

import (
	"Food-Delivery/internal/order/entity/order_model"
	user_model "Food-Delivery/internal/user/entity/model"
	"Food-Delivery/pkg/common"
)

type MenuItem struct {
	Id           int `json:"id" gorm:"column:id;"`
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id;"`
	//Name         string        `json:"name" gorm:"column:name;varchar(255)"`
	Description *string       `json:"description" gorm:"column:description;"`
	Price       float64       `json:"price" gorm:"column:price;"`
	Image       *common.Image `json:"image" gorm:"column:image"`
}

func (MenuItem) TableName() string {
	return "menu_item"
}

//===========================================================================================

type Order struct {
	Id           int                     `json:"id" gorm:"column:id;"`
	UserId       int                     `json:"menu_id" gorm:"column:menu_id;"`
	User         user_model.User         `json:"user" gorm:"foreignKey:UserId;references:Id;"`
	RestaurantId int                     `json:"-" gorm:"column:restaurant_id;"`
	OrderItem    []OrderItem             `json:"order_items" gorm:"foreignKey:OrderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalAmount  float64                 `json:"total_amount" gorm:"column:total_amount;"`
	Status       order_model.OrderStatus `json:"status" gorm:"column:status;"`
}

type OrderItem struct {
	Id         int     `json:"id" gorm:"column:id;"`
	OrderId    int     `json:"order_id" gorm:"column:order_id;"`
	MenuItemId int     `json:"menu_item_id" gorm:"column:menu_item_id;"`
	Quantity   int     `json:"quantity" gorm:"column:quanity;"`
	Note       *string `json:"note" gorm:"column:note;"`
}

func (Order) TableName() string {
	return "order"
}

func (OrderItem) TableName() string {
	return "order_item"
}

// ===========================================================================================

type Category struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Name   string `json:"name" gorm:"column:name;"`
	Status string `json:"status" gorm:"column:status;"`
	//RestaurantId int    `json:"restaurant_id" gorm:"column:restaurant_id;"`
}

func (Category) TableName() string {
	return "category"
}
