package driver_dto

import (
	"Food-Delivery/pkg/common"
)

type CreateDTO struct {
	Email     string        `json:"email" gorm:"column:email"`
	FirstName string        `json:"firstName" gorm:"column:first_name"`
	LastName  string        `json:"lastName" gorm:"column:last_name"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (dto *CreateDTO) Validate() error {

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
