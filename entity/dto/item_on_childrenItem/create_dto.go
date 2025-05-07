package item_on_childrenItem_dto

import "Food-Delivery/pkg/common"

type CreateDTO struct {
	RestaurantId     int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId       int           `json:"category_id" gorm:"column:category_id;"`
	VendorCategoryId int           `json:"vendor_category_id" gorm:"column:vendor_category_id;"`
	Name             string        `json:"name" gorm:"column:name;"`
	Image            *common.Image `json:"image" gorm:"column:image;"`
	Price            float32       `json:"price" gorm:"column:price;"`
	Description      *string       `json:"description" gorm:"column:description;"`
}
