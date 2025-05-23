package item_dto

import rating_dto "Food-Delivery/entity/dto/rating"

type ItemDTO struct {
	ID           int                   `json:"id"`
	Name         string                `json:"name"`
	Price        float64               `json:"price"`
	Description  *string               `json:"description,omitempty"`
	DeliveryTime int                   `json:"delivery_time" `
	CategoryId   int                   `json:"category_id"`
	RestaurantId int                   `json:"restaurant_id,omitempty"`
	Rating       *rating_dto.RatingDTO `json:"rating,omitempty"`
}
