package model

import (
	"Food-Delivery/pkg/common"
)

const RatingEntity = "rating"

type Rating struct {
	common.SQLModel
	UserId       int    `json:"user_id" gorm:"column:user_id;"`
	RestaurantId *int   `json:"restaurant_id" gorm:"column:restaurant_id;"`
	ItemId       *int   `json:"item_id" gorm:"column:item_id;"`
	Like         bool   `json:"like" gorm:"column:like;default:true"`
	Score        int    `json:"score" gorm:"column:score;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
}

func (rating Rating) TableName() string {
	return "rating"
}

func (r *Rating) ToData() map[string]interface{} {

	data := map[string]interface{}{
		"id":     r.Id,
		"userId": r.UserId,
	}

	if r.RestaurantId != nil {
		data["restaurant_id"] = *r.RestaurantId
	}

	if r.ItemId != nil {
		data["item_id"] = *r.ItemId
	}

	return data
}
