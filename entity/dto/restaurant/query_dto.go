package restaurant_dto

import "Food-Delivery/entity/constant"

type QueryDTO struct {
	Status    *constant.RestaurantStatus `form:"status"`
	SearchKey *string                    `form:"search_key"`
}
