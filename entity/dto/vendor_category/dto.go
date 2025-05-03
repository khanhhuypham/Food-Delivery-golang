package vendor_category_dto

import (
	rating_dto "Food-Delivery/entity/dto/rating"
	"Food-Delivery/pkg/common"
)

type VendorCategoryDTO struct {
	ID           int           `json:"id"`
	Image        *common.Image `json:"image"`
	Name         string        `json:"name"`
	Description  *string       `json:"description"`
	Active       bool          `json:"active"`
	RestaurantId int           `json:"restaurant_id"`
	TotalItems   int           `json:"total_items,omitempty"`
	Items        []ItemDTO     `json:"items,omitempty"`
}

type ItemDTO struct {
	ID               int                   `json:"id"`
	Name             string                `json:"name"`
	Image            *common.Image         `json:"image"`
	Price            float64               `json:"price"`
	Description      *string               `json:"description,omitempty"`
	DeliveryTime     int                   `json:"delivery_time" `
	CategoryId       int                   `json:"category_id"`
	VendorCategoryId int                   `json:"vendor_category_id"`
	RestaurantId     int                   `json:"restaurant_id"`
	Rating           *rating_dto.RatingDTO `json:"rating,omitempty"`
}
