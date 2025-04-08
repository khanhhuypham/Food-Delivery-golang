package usermodel

import (
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"errors"
	"strings"
)

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (dto *UserLogin) Validate() error {
	dto.Email = strings.TrimSpace(dto.Email)
	dto.Password = strings.TrimSpace(dto.Password)

	if !utils.CheckValidEmailFormat(dto.Email) {
		return ErrInvalidEmail
	}

	//if len(dto.Password) < 6 {
	//	return ErrInvalidPassword
	//}

	return nil

}

type UserCreate struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"password" gorm:"column:password"`
	FirstName string        `json:"firstName" gorm:"column:first_name"`
	LastName  string        `json:"lastName" gorm:"column:last_name"`
	Role      string        `json:"-" form:"column:role"` /*???????? why form*/
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (dto *UserCreate) PrepareCreate() error {
	//HashPassword
	hashedPwd, err := utils.HashPassword(dto.Password)
	if err != nil {
		return err
	}
	dto.Password = hashedPwd

	//set Default role
	dto.Role = "guest"
	return nil
}

func (dto *UserCreate) Validate() error {
	dto.Email = strings.TrimSpace(dto.Email)
	dto.Password = strings.TrimSpace(dto.Password)
	dto.FirstName = strings.TrimSpace(dto.FirstName)
	dto.LastName = strings.TrimSpace(dto.LastName)

	if dto.Email == "" {
		return errors.New("email can't be blank")
	}

	if !utils.CheckValidEmailFormat(dto.Email) {
		return ErrInvalidEmail
	}

	if len(dto.Password) < 6 {
		return ErrInvalidPassword
	}

	if dto.FirstName == "" {
		return ErrInvalidFirstName
	}

	if dto.LastName == "" {
		return ErrInvalidLastName
	}

	//another validation for password will need to write below

	return nil
}
