package model

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

const CategoryEntity = "category"

type Category struct {
	common.SQLModel
	Image       *common.Image           `json:"image" gorm:"column:image;"`
	Name        string                  `json:"name" gorm:"column:name;not null;unique"`
	Description *string                 `json:"description" gorm:"column:description;"`
	Active      bool                    `json:"active" gorm:"column:active;default:true"`
	Status      constant.CategoryStatus `json:"status" gorm:"column:status;"`
	Items       []Item                  `json:"items" gorm:"foreignKey:CategoryId;references:Id"`
}

func (cate *Category) TableName() string {
	return "category"
}

func (cate *Category) Validate() error {
	cate.Name = strings.TrimSpace(cate.Name)

	if cate.Name == "" {
		return common.ErrBadRequest(errors.New("category name is required"))
	}

	if !cate.Status.IsValid() {
		return common.ErrBadRequest(errors.New("invalid category status"))
	}

	return nil
}
