package children_item_service

import (
	children_item_dto "Food-Delivery/entity/dto/children_item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
)

type childrenItemRepository interface {
	FindAllWithCondition(ctx context.Context, query *children_item_dto.QueryDTO, keys ...string) ([]model.ChildrenItem, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.ChildrenItem, error)
	Create(ctx context.Context, dto *children_item_dto.CreateDTO) (*model.ChildrenItem, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *children_item_dto.CreateDTO) (*model.ChildrenItem, error)
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

type childrenItemService struct {
	childrenItemRepo childrenItemRepository
}

func NewChildrenItemService(childrenItemRepo childrenItemRepository) *childrenItemService {
	return &childrenItemService{childrenItemRepo}
}

func (service *childrenItemService) FindAll(ctx context.Context, filter *children_item_dto.QueryDTO) ([]model.ChildrenItem, error) {
	//there will have business logic before getting data list with condition
	childrenItems, err := service.childrenItemRepo.FindAllWithCondition(ctx, filter)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return childrenItems, nil
}

func (service *childrenItemService) FindOneById(ctx context.Context, id int) (*model.ChildrenItem, error) {
	//there will have business logic before getting specific data with condition

	childrenItem, err := service.childrenItemRepo.FindOneWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return childrenItem, nil
}

func (service *childrenItemService) Create(ctx context.Context, dto *children_item_dto.CreateDTO) (*model.ChildrenItem, error) {
	//------perform business operation such as validate data
	if err := dto.Validate(); err != nil {
		return nil, common.ErrBadRequest(err)
	}

	newData, err := service.childrenItemRepo.Create(ctx, dto)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return newData, nil
}

func (service *childrenItemService) Update(ctx context.Context, id int, dto *children_item_dto.CreateDTO) (*model.ChildrenItem, error) {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	//check the eixstence of data in database
	if _, err := service.FindOneById(ctx, id); err != nil {
		return nil, err
	}

	updatedData, err := service.childrenItemRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return updatedData, nil
}

func (service *childrenItemService) Delete(ctx context.Context, id int) error {
	//check the existence of data in database
	if _, err := service.FindOneById(ctx, id); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.childrenItemRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
