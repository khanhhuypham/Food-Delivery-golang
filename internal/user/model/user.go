package usermodel

import (
	"Food-Delivery/pkg/common"
)

const EntityName = "user"

type User struct {
	common.SQLModel
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"-" gorm:"column:password"`
	FirstName string        `json:"firstName" gorm:"column:first_name"`
	LastName  string        `json:"lastName" gorm:"column:last_name"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Role      UserRole      `json:"role" gorm:"column:role"`
	Status    UserStatus    `json:"status" gorm:"column:status;"`
}

type UserType string
type UserRole string
type UserStatus string

const (
	TypeEmailPassword UserType = "email_password"
	TypeFacebook      UserType = "facebook"
	TypeGmail         UserType = "gmail"

	RoleUser    UserRole = "user"
	RoleAdmin   UserRole = "admin"
	RoleShipper UserRole = "shipper"

	StatusPending  UserStatus = "pending"
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
	StatusBanned   UserStatus = "banned"
	StatusDeleted  UserStatus = "deleted"
)

func (User) TableName() string {
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
