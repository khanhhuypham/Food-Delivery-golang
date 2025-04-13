package order_dto

type QueryDTO struct {
	SearchKey *string `json:"search_key" form:"search_key"`
}
