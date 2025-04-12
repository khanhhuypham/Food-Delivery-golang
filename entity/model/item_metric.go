package model

import (
	"Food-Delivery/pkg/common"
	"time"
)

type ItemMetric struct {
	common.SQLModel
	ItemID       int       `json:"item_id" gorm:"column:item_id;"`
	TotalOrders  int       `json:"total_orders" gorm:"column:total_orders;"`
	WeeklyOrders int       `json:"weekly_orders" gorm:"column:weekly_orders;"`
	AvgRating    float64   `json:"avg_rating" gorm:"column:avg_rating;"`
	TotalReviews int       `json:"total_reviews" gorm:"column:total_reviews;"`
	LastOrderAt  time.Time `json:"last_order_at" gorm:"column:last_order_at;"`
}

func (metric *ItemMetric) TableName() string {
	return "item_metric"
}

/*
2 Item Metrics Table
* Purpose: Aggregates performance metrics for menu items,
such as average rating, order count, and view count, to simplify queries and improve performance.
* Structure (from previous response): sql CollapseWrapCopy

* Why Use It:
    * Reduces the need for expensive joins or aggregations across ratings, order_items, or user_interactions.
    * Provides a snapshot of item performance for rankings like "best food" (high avg_rating) or "most popular" (high order_count).
    * Supports business analytics, e.g., identifying underperforming items.
*/
