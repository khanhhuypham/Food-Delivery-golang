package relationship

import (
	menu_item_model "Food-Delivery/internal/menu_item/model"
)

type OrderItem struct {
	Id         int                      `gorm:"column:id;"`
	OrderId    int                      `json:"order_id" gorm:"column:order_id;"`
	MenuItemId int                      `json:"menu_item_id" gorm:"column:menu_item_id;"`
	MenuItem   menu_item_model.MenuItem `json:"menu_item" gorm:"foreignKey:MenuItemId;references:Id;"`
	Quantity   int                      `json:"quantity" gorm:"column:quanity;"`
	Note       *string                  `json:"note" gorm:"column:note;"`
}
