package model

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
)

/*
3 User Interactions Table
	* Purpose: Tracks user actions like views, likes, or shares to capture engagement with menu items,
	which feeds into popularity and recommendation algorithms.


* Why Use It:
    * Captures fine-grained user behavior beyond orders (e.g., browsing patterns).
    * Helps identify trending items based on recent interactions.
    * Feeds data into machine learning models for personalized recommendations.
*/

type UserInteraction struct {
	common.SQLModel
	UserID int                       `gorm:"column:user_id;"`
	ItemID int                       `gorm:"column:item_id;"`
	Score  int                       `gorm:"column:score;"`
	Type   constant.Interaction_type `gorm:"column:interaction_type;"`
}
