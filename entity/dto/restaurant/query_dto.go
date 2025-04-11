package restaurant_dto

type QueryDTO struct {
	OwnerId *string `json:"ownerId"`
	CityId  *int    `json:"cityId"`
	Status  *string `json:"status"`
}
