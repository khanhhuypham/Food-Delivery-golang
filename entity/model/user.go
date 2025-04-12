package model

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
)

const UserEntity = "user"

type User struct {
	common.SQLModel
	Avatar    *common.Image       `json:"avatar" gorm:"column:avatar"`
	Email     string              `json:"email" gorm:"column:email"`
	Password  string              `json:"-" gorm:"column:password"`
	FirstName string              `json:"firstName" gorm:"column:first_name"`
	LastName  string              `json:"lastName" gorm:"column:last_name"`
	Phone     string              `json:"phone" gorm:"column:phone"`
	Role      constant.UserRole   `json:"role" gorm:"column:role"`
	Status    constant.UserStatus `json:"status" gorm:"column:status;"`
	Orders    []Order             `gorm:"foreignKey:UserId"`
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) GetUserEmail() string {
	return user.Email
}

func (user *User) GetFullName() string {
	return user.FirstName + " " + user.LastName
}

func (user *User) GetUserRole() string {
	return string(user.Role)
}
