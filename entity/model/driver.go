package model

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
)

const DriverEntity = "driver"

type Driver struct {
	common.SQLModel
	IdCardNumber   string              `json:"id_card_number" gorm:"column:id_card_number"`
	Avatar         *common.Image       `json:"avatar" gorm:"column:avatar;default:null"`
	Insurance      *common.Image       `json:"insurance" gorm:"column:insurance"`
	DrivingLicense *common.Image       `json:"driving_license" gorm:"column:driving_license"`
	Email          string              `json:"email" gorm:"column:email"`
	Name           string              `json:"name" gorm:"column:name;not null;type:varchar(255);uniqueIndex:idx_driver_name"`
	Phone          string              `json:"phone" gorm:"column:phone"`
	Address        string              `json:"address" gorm:"column:address"`
	Status         constant.UserStatus `json:"status" gorm:"column:status;"`
	Orders         []Order             `gorm:"foreignKey:DriverId;references:Id;constraint:OnUpdate:CASCADE"`
	//Ratings   []Rating            `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (driver Driver) TableName() string {
	return "driver"
}
