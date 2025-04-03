package categorymodel

import (
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

type QueryDTO struct {
	Status string `json:"status"`
}
type CategoryCreateDto struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

func (CategoryCreateDto) TableName() string {
	return Category{}.TableName()
}

func (dto *CategoryCreateDto) Validate() error {
	if status := dto.Status; status != nil && *status != StatusActive && *status != StatusInactive && *status != StatusDeleted {
		return common.ErrBadRequest(errors.New("status is invalid"))
	}

	if str := dto.Name; str != nil {

		*dto.Name = strings.TrimSpace(*str)

		if len(*dto.Name) == 0 {
			return common.ErrBadRequest(errors.New("name is required"))
		}
	}

	return nil
}
