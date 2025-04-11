package menu_item_service

import (
	menu_item_dto "Food-Delivery/entity/dto/item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(ctx context.Context, dto *menu_item_dto.CreateDTO) (*model.Item, error)
	FindAllWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *menu_item_dto.QueryDTO,
		keys ...string) ([]model.Item, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Item, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *menu_item_dto.CreateDTO) (*model.Item, error)
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

type itemService struct {
	itemRepo ItemRepository
}

func NewRestaurantService(itemRepo ItemRepository) *itemService {
	return &itemService{itemRepo}
}

func (service *itemService) Create(ctx context.Context, dto *menu_item_dto.CreateDTO) (*model.Item, error) {
	//------perform business operation such as validate data
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	newItem, err := service.itemRepo.Create(ctx, dto)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return newItem, nil
}

func (service *itemService) FindAll(ctx context.Context, paging *common.Paging, query *menu_item_dto.QueryDTO) ([]model.Item, error) {
	//there will have business logic before getting data list with condition
	items, err := service.itemRepo.FindAllWithCondition(ctx, paging, query)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return items, nil
}

func (service *itemService) FindOneById(ctx context.Context, id int) (*model.Item, error) {
	//there will have business logic before getting specific data with condition

	item, err := service.itemRepo.FindOneWithCondition(ctx, map[string]any{"id": id}, "Restaurant")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.ItemEntity, err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return item, nil
}

func (service *itemService) Update(ctx context.Context, id int, dto *menu_item_dto.CreateDTO) (*model.Item, error) {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	//check the eixstence of data in database
	if _, err := service.itemRepo.FindOneWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return nil, err
	}

	updatedItem, err := service.itemRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return updatedItem, nil
}

func (service *itemService) Delete(ctx context.Context, id int) error {
	//check the eixstence of data in database
	if _, err := service.itemRepo.FindOneWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.itemRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
