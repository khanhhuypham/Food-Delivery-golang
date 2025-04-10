package order_model

type OrderCreateDTO struct {
	UserId       int     `json:"user_id"`
	RestaurantId int     `json:"restaurant_id"`
	Description  *string `json:"description"`
}

func (dto *OrderCreateDTO) Validate() error {

	//dto.Name = strings.TrimSpace(dto.Name)
	//
	//if len(dto.Name) == 0 {
	//	return common.ErrBadRequest(errors.New("restaurant name is empty"))
	//}
	//
	//if dto.Price <= 0 {
	//	return common.ErrBadRequest(errors.New("price must be greater than zero"))
	//}

	return nil
}

type OrderUpdateDTO struct {
	Status OrderStatus `json:"status" form:"status"`
}

func (dto *OrderUpdateDTO) Validate() error {

	//dto.Name = strings.TrimSpace(dto.Name)
	//
	//if len(dto.Name) == 0 {
	//	return common.ErrBadRequest(errors.New("restaurant name is empty"))
	//}
	//
	//if dto.Price <= 0 {
	//	return common.ErrBadRequest(errors.New("price must be greater than zero"))
	//}

	return nil
}
