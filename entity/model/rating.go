package model

import (
	"Food-Delivery/pkg/common"
)

type Rating struct {
	common.SQLModel
	UserId       int    `json:"user_id" gorm:"column:user_id;not null"`
	RestaurantId int    `json:"restaurant_id" gorm:"column:restaurant_id; not null"`
	ItemId       *int   `json:"item_id" gorm:"column:item_id;"`
	Score        int    `json:"score" gorm:"column:score;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
}

func (rating *Rating) TableName() string {
	return "rating"
}
