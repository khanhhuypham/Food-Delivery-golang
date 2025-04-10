package order_item_model

import (
	menu_item_model "Food-Delivery/internal/menu_item/model"
	"Food-Delivery/pkg/common"
)

const EntityName = "order item"

type OrderItem struct {
	common.SQLModel
	OrderId int `json:"order_id" gorm:"column:order_id;"`
	//Order      order_model.Order        `json:"order" gorm:"foreignKey:OrderId;references:Id;"`
	MenuItemId int                      `json:"menu_item_id" gorm:"column:menu_item_id;"`
	MenuItem   menu_item_model.MenuItem `json:"menu_item" gorm:"foreignKey:MenuItemId;references:Id;"`
	Quantity   int                      `json:"quantity" gorm:"column:quanity;"`
	Note       *string                  `json:"note" gorm:"column:note;"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
