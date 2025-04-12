package category_dto

import (
	"Food-Delivery/entity/constant"
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

// ======================================= query dto ========================================
type QueryDTO struct {
	Active *bool `json:"active"`
}

// ======================================= query dto ========================================
type RPCRequestDTO struct {
	Ids []int `json:"ids"`
}

// ======================================= create dto ========================================
type CreateDto struct {
	Name        *string                  `json:"name"`
	Description *string                  `json:"description"`
	Status      *constant.CategoryStatus `json:"status"`
}

func (dto *CreateDto) Validate() error {
	if status := dto.Status; status != nil && !status.IsValid() {
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
