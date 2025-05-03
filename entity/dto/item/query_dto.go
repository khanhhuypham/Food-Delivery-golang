package item_dto

type QueryDTO struct {
	SearchKey       *string `form:"search_key"`
	DescendingPrice *bool   `form:"descending_price"`
	PriceRange      *bool   `form:"price_range"`
	DescendingScore *bool   `form:"descending_score"`
}
