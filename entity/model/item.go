package model

import (
	item_dto "Food-Delivery/entity/dto/item"
	rating_dto "Food-Delivery/entity/dto/rating"
	"Food-Delivery/pkg/common"
)

const ItemEntity = "menu item"

type Item struct {
	common.SQLModel
	Image            *common.Image `json:"image" gorm:"column:image;"`
	RestaurantId     int           `json:"restaurant_id" gorm:"column:restaurant_id;not null"`
	Name             string        `json:"name" gorm:"column:name;not null"`
	Description      *string       `json:"description" gorm:"column:description;"`
	Price            float64       `json:"price" gorm:"column:price;not null"`
	DeliveryTime     int           `json:"delivery_time" gorm:"column:delivery_time;not null;default:0"`
	Rating           *Rating       `json:"rating" gorm:"foreignKey:ItemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Optionals        []Optional    `json:"optionals" gorm:"foreignKey:ItemId;references:Id"`
	CategoryId       int           `json:"category_id" gorm:"column:category_id;not null"`
	VendorCategoryId int           `json:"vendor_category_id" gorm:"column:vendor_category_id;not null"`

	//OrderItems       []OrderItem           `json:"order_items" gorm:"foreignKey:ItemId;references:Id"`
	//ChildrenItems    []ItemOnChildrenItems `json:"children_items" gorm:"foreignKey:ItemId;references:Id"`
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
		DeliveryTime: item.DeliveryTime,
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

func (item *Item) ToItemDetailDTO() *item_dto.ItemDTO {
	dto := &item_dto.ItemDTO{
		ID:          item.Id,
		Name:        item.Name,
		Price:       item.Price,
		Description: item.Description,
		CategoryId:  item.CategoryId,
	}

	if item.Rating != nil {
		dto.Rating = &rating_dto.RatingDTO{
			Like:  item.Rating.Like,
			Score: item.Rating.Score,
		}
	}

	return dto
}
