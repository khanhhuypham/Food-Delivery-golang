package categorymodel

import (
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

const EntityName = "category"

type Category struct {
	common.SQLModel
	Name        string `json:"name" gorm:"column:name;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      Status `json:"status" gorm:"column:status;"`
	//CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;"`
	//UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updated_at;"`
}

func (Category) TableName() string {
	return "category"
}

// Define Status as a custom type
type Status string

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusDeleted  = "deleted"
)

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

// Method to check if the status is valid
func (s Status) IsValid() bool {
	switch s {
	case StatusActive, StatusInactive, StatusDeleted:
		return true
	}
	return false
}
