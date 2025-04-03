package usermodel

import (
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"errors"
)

const EntityName = "user"

type User struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"-" gorm:"column:password"`
	FirstName string        `json:"firstName" gorm:"column:first_name"`
	LastName  string        `json:"lastName" gorm:"column:last_name"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Role      string        `json:"role" gorm:"column:role"`
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) GetUserEmail() string {
	return user.Email
}

func (user *User) GetUserRole() string {
	return user.Role
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
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

func (userCreate *UserCreate) PrepareCreate() error {
	//HashPassword
	hashedPwd, err := utils.HashPassword(userCreate.Password)
	if err != nil {
		return err
	}
	userCreate.Password = hashedPwd

	//set Default role
	userCreate.Role = "guest"
	return nil
}

func (userCreate *UserCreate) Validate() error {
	if userCreate.Email == "" {
		return errors.New("email can't be blank")
	}
	//another validation for password will need to write below

	return nil
}
