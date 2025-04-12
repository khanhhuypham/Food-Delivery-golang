package model

import (
	"gorm.io/gorm"
	"time"
)

const OrderItemEntity = "order item"

type OrderItem struct {
	OrderId int `json:"order_id" gorm:"primaryKey;column:order_id"`
	ItemId  int `json:"item_id" gorm:"primaryKey;column:item_id"`

	// Associations
	Order *Order `gorm:"foreignKey:OrderId;references:Id"`
	Item  *Item  `gorm:"foreignKey:ItemId;references:Id"`

	Quantity  int            `json:"quantity" gorm:"column:quantity"`
	Note      *string        `json:"note" gorm:"column:note"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

func (orderItem *OrderItem) TableName() string {
	return "order_item"
}
