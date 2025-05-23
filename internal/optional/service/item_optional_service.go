package item_optional_service

import (
	item_optional_dto "Food-Delivery/entity/dto/item_optional"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type itemOptionalRepository interface {
	Create(ctx context.Context, dto *item_optional_dto.CreateDTO) (*model.Optional, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *item_optional_dto.CreateDTO) (*model.Optional, error)
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
	FindAllWithCondition(ctx context.Context, restaurantId int, keys ...string) ([]model.Optional, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Optional, error)

	//========================================Children item=================
	FindChildrenItemByIds(ctx context.Context, ids []int) ([]model.ChildrenItem, error)
	RemoveChildrenItemFromOptional(ctx context.Context, optional *model.Optional) (*model.Optional, error)
	AddChildrenItemToOptional(ctx context.Context, optionalId int, childrenItemIds []int) (*model.Optional, error)
}

type itemOptionalService struct {
	repo itemOptionalRepository
}

func NewItemOptionalService(repo itemOptionalRepository) *itemOptionalService {
	return &itemOptionalService{repo}
}

func (service *itemOptionalService) Create(ctx context.Context, dto *item_optional_dto.CreateDTO) (*model.Optional, error) {
	//------perform business operation such as validate data
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	newItem, err := service.repo.Create(ctx, dto)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	if _, err := service.AddChildrenItemToOptional(ctx, newItem.Id, dto.ChildrenItemId); err != nil {
		return nil, err
	}

	//Find UpdateItem again after updated and add children Item
	newItem, err = service.repo.FindOneWithCondition(ctx, map[string]any{"id": newItem.Id}, "ChildrenItems")

	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func (service *itemOptionalService) Update(ctx context.Context, id int, dto *item_optional_dto.CreateDTO) (*model.Optional, error) {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	//check the eixstence of data in database
	if _, err := service.repo.FindOneWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return nil, err
	}

	updatedItem, err := service.repo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto)

	if _, err := service.AddChildrenItemToOptional(ctx, id, dto.ChildrenItemId); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	//Find UpdateItem again after updated and add children Item
	updatedItem, err = service.repo.FindOneWithCondition(ctx, map[string]any{"id": id}, "ChildrenItems")

	if err != nil {
		return nil, err
	}

	return updatedItem, nil
}

func (service *itemOptionalService) Delete(ctx context.Context, id int) error {
	//check the eixstence of data in database
	if _, err := service.repo.FindOneWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.repo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

//=========================================Query=========================================

func (service *itemOptionalService) FindAll(ctx context.Context, restaurantId int) ([]model.Optional, error) {
	//there will have business logic before getting data list with condition
	data, err := service.repo.FindAllWithCondition(ctx, restaurantId, "ChildrenItems")

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return data, nil
}

func (service *itemOptionalService) FindOneById(ctx context.Context, id int) (*model.Optional, error) {
	//there will have business logic before getting specific data with condition

	item, err := service.repo.FindOneWithCondition(ctx, map[string]any{"id": id}, "ChildrenItems")

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.ItemOptionalEntity, err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return item, nil
}
