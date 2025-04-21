package rating_dto

import (
	"Food-Delivery/pkg/common"
	"errors"
)

// DTO = Data Transfer Object
type CreateDTO struct {
	UserId       int     `json:"user_id" form:"user_id" gorm:"column:user_id;"`
	RestaurantId *int    `json:"restaurant_id" form:"restaurant_id" gorm:"column:restaurant_id;"`
	ItemId       *int    `json:"item_id" form:"item_id" gorm:"column:item_id;"`
	Like         *bool   `json:"like" form:"like" gorm:"column:like;"`
	Score        *int    `json:"score" gorm:"column:score;"`
	Comment      *string `json:"comment" gorm:"column:comment;"`
}

func (dto *CreateDTO) Validate() error {
	restaurantId := dto.RestaurantId
	ItemId := dto.ItemId

	if ItemId != nil && restaurantId != nil {
		return common.ErrBadRequest(errors.New("User can only like for one of 2 two entities, restaurant or item"))
	} else if ItemId == nil && restaurantId == nil {
		return common.ErrBadRequest(errors.New("restaurant_id or item_id must be required"))
	}

	if score := dto.Score; score != nil && *dto.Score < 0 {
		return common.ErrBadRequest(errors.New("score must greater than 0"))
	}

	return nil
}

func (dto *CreateDTO) ToData() map[string]interface{} {

	data := map[string]interface{}{
		"user_id": dto.UserId,
		"like":    dto.Like,
	}

	if dto.RestaurantId != nil {
		data["restaurant_id"] = *dto.RestaurantId
	}

	if dto.ItemId != nil {
		data["item_id"] = *dto.ItemId
	}

	return data
}
