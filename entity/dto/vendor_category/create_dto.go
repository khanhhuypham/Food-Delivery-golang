package vendor_category_dto

import "Food-Delivery/pkg/common"

type CreateDTO struct {
	Name         *string       `json:"name" gorm:"column:name;"`
	Image        *common.Image `json:"image" gorm:"column:image;"`
	RestaurantId int           `json:"restaurant_id" gorm:"column:restaurant_id"`
	Description  *string       `json:"description" gorm:"column:description;"`
}

func (dto *CreateDTO) Validate() error {
	//if status := dto.Status; status != nil && !dto.Status.IsValid() {
	//	return common.ErrBadRequest(errors.New("status invalid"))
	//}
	//
	//if str := dto.Name; str != nil {
	//	*dto.Name = strings.TrimSpace(*str)
	//
	//	if len(*dto.Name) == 0 {
	//		return common.ErrBadRequest(errors.New("restaurant name is empty"))
	//	}
	//}
	//
	//if addr := dto.Address; addr != nil {
	//	*dto.Address = strings.TrimSpace(*addr)
	//
	//	if len(*dto.Address) == 0 {
	//		return common.ErrBadRequest(errors.New("restaurant address is required"))
	//	}
	//}

	return nil
}
