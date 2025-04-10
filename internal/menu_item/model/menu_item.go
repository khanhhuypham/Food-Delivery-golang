package menu_item_model

import (
	"Food-Delivery/pkg/common"
)

const EntityName = "menu_item"

type MenuItem struct {
	common.SQLModel
	RestaurantId int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Restaurant   Restaurant    `json:"restaurant" gorm:"foreignKey:RestaurantId;references:Id;"`
	Name         string        `json:"name" gorm:"column:name;type:varchar(255);uniqueIndex:idx_menu_item_name"`
	Description  *string       `json:"description" gorm:"column:description;"`
	Price        float64       `json:"price" gorm:"column:price;"`
	Image        *common.Image `json:"image" gorm:"column:image"`
}

func (MenuItem) TableName() string {
	return "menu_item"
}
