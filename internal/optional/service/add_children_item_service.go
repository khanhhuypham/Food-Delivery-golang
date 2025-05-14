package item_optional_service

import (
	"Food-Delivery/entity/model"
	"context"
)

func (service *itemOptionalService) AddChildrenItemToOptional(ctx context.Context, optionalId int, childrenItemId []int) (*model.Optional, error) {
	var result *model.Optional

	// 1️⃣ Check if the children items exist in the database
	optional, err := service.FindOneById(ctx, optionalId)

	if err != nil {
		return nil, err
	}

	result, err = service.repo.RemoveChildrenItemFromOptional(ctx, optional)
	if err != nil {
		return nil, err
	}

	result, err = service.repo.AddChildrenItemToOptional(ctx, optionalId, childrenItemId)
	if err != nil {
		return nil, err
	}

	// 🎉 If everything goes well, return nil (no errors)
	return result, nil
}
