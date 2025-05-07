package model

import "Food-Delivery/pkg/common"

type ItemOnChildrenItems struct {
	common.SQLModel
	ItemId         int           `json:"item_id" gorm:"column:item_id;not null"`
	Item           *Item         `gorm:"foreignKey:ItemId;references:Id"`
	ChildrenItemId int           `json:"order_id" gorm:"column:order_id;not null"`
	ChildrenItem   *ChildrenItem `gorm:"foreignKey:ChildrenItemId;references:Id"`
	Quantity       int           `json:"quantity" gorm:"column:quantity;not null"`
	Note           *string       `json:"note" gorm:"column:note"`
}

func (ItemOnChildrenItems) TableName() string {
	return "item_on_children_item"
}
