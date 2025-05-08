package item_optional_dto

type CreateDTO struct {
	RestaurantId   int     `json:"restaurant_id" gorm:"column:restaurant_id;"`
	ItemId         int     `json:"item_id" gorm:"column:item_id;"`
	Name           string  `json:"name" gorm:"column:name;"`
	Description    *string `json:"description" gorm:"column:description;"`
	ChildrenItemId []int   `json:"children_items"`
}

func (dto *CreateDTO) Validate() error {
	return nil
}
