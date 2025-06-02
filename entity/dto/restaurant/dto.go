package restaurant_dto

import (
	"Food-Delivery/entity/constant"
	rating_dto "Food-Delivery/entity/dto/rating"
	"Food-Delivery/pkg/common"
)

type RestaurantDTO struct {
	ID          int                       `json:"id"`
	Cover       *common.Image             `json:"cover"`
	Logo        *common.Image             `json:"logo"`
	Name        string                    `json:"name"`
	LikeCount   int                       `json:"like_count" `
	Status      constant.RestaurantStatus `json:"status"`
	Description *string                   `json:"description"`
	Rating      *rating_dto.RatingDTO     `json:"rating,omitempty"`
}

type RestaurantDetailDTO struct {
	ID          int                       `json:"id"`
	Name        string                    `json:"name"`
	Description *string                   `json:"description,omitempty"`
	Status      constant.RestaurantStatus `json:"status"`
	Address     string                    `json:"address"`
	Rating      *rating_dto.RatingDTO     `json:"rating,omitempty"`
}
