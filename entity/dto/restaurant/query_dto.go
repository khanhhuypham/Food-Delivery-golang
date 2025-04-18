package restaurant_dto

type QueryDTO struct {
	Active    *bool   `form:"active"`
	Status    *string `form:"status"`
	SearchKey *string `form:"search_key"`
}
