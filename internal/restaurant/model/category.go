package restaurant_model

import "github.com/google/uuid"

type Category struct {
	Id     uuid.UUID `json:"id" gorm:"column:id;"`
	Name   string    `json:"name" gorm:"column:name;"`
	Status string    `json:"status" gorm:"column:status;"`
}

func (Category) TableName() string {
	return "category"
}
