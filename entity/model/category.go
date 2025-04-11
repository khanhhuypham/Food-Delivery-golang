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
	Name        string                  `json:"name" gorm:"column:name;"`
	Description string                  `json:"description" gorm:"column:description;"`
	Status      constant.CategoryStatus `json:"status" gorm:"column:status;"`
}

func (Category) TableName() string {
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
