package children_item_dto

type QueryDTO struct {
	SearchKey    *string `form:"search_key"`
	RestaurantId int     `form:"restaurant_id"`
}
