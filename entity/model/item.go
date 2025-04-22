package model

import (
	item_dto "Food-Delivery/entity/dto/item"
	rating_dto "Food-Delivery/entity/dto/rating"
	"Food-Delivery/pkg/common"
)

const ItemEntity = "menu item"

type Item struct {
	common.SQLModel
	Image        *Media       `json:"image" gorm:"column:image;"`
	RestaurantId int          `json:"restaurant_id" gorm:"column:restaurant_id;not null"`
	Name         string       `json:"name" gorm:"column:name;type:varchar(255);not null;uniqueIndex:idx_item_name"`
	Description  *string      `json:"description" gorm:"column:description;"`
	Price        float64      `json:"price" gorm:"column:price;not null"`
	OrderItems   []*OrderItem `json:"orderItems" gorm:"foreignKey:ItemId;references:Id"`
	CategoryId   int          `json:"category_id" gorm:"column:category_id;not null"`
	Rating       *Rating      `json:"rating" gorm:"foreignKey:ItemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (item Item) TableName() string {
	return "item"
}

func (item *Item) ToItemDTO() *item_dto.ItemDTO {
	dto := &item_dto.ItemDTO{
		ID:           item.Id,
		Name:         item.Name,
		Price:        item.Price,
		Description:  item.Description,
		CategoryId:   item.CategoryId,
		RestaurantId: item.RestaurantId,
	}

	if item.Rating != nil {
		dto.Rating = &rating_dto.RatingDTO{
			Like:  item.Rating.Like,
			Score: item.Rating.Score,
		}
	}

	return dto
}
