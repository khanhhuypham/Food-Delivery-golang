package model

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
)

/*
	1 Recommendations Table
* Purpose:
	Stores precomputed or dynamically generated recommendations for users,
	such as "best food," "most recommended," or "most popular."

	It acts as a centralized place to serve personalized or trending items quickly.

* Why Use It:
    * Enables fast retrieval of curated lists (e.g., top 10 popular dishes) without real-time computation.
    * Supports personalized recommendations by linking to user_id.
    * Facilitates A/B testing of recommendation algorithms by storing different recommendation_type values.

* Leverage Strategy:
	Use the user_interactions, orders, or rating table to feed recommendation algorithm (e.g, collaborative filltering)
	- Advanced option:
		* Use a machine learning model (e.g., matrix factorization) trained on user_interactions and ratings
		* Store predictions in recommendation with score as the confidence level
*/

type Recommend struct {
	common.SQLModel
	UserID int                          `gorm:"column:user_id;"`
	ItemID int                          `gorm:"column:item_id;"`
	Score  int                          `gorm:"column:score;"`
	Type   constant.Recommendation_Type `gorm:"column:recommendation_type;"`
}
