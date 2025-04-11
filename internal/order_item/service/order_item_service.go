package order_item_service

import (
	order_item_dto "Food-Delivery/entity/dto/order-item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(ctx context.Context, dto *order_item_dto.CreateDTO) error
	FindAllWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *order_item_dto.QueryDTO,
		keys ...string) ([]model.OrderItem, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.OrderItem, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *order_item_dto.CreateDTO) error
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

type orderItemService struct {
	orderItemRepo OrderItemRepository
}

func NewOrderItemService(orderItemRepo OrderItemRepository) *orderItemService {
	return &orderItemService{orderItemRepo}
}

func (service *orderItemService) Create(ctx context.Context, dto *order_item_dto.CreateDTO) error {
	//------perform business operation such as validate data
	if err := dto.Validate(); err != nil {
		return err
	}
	//------
	if err := service.orderItemRepo.Create(ctx, dto); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *orderItemService) FindAll(ctx context.Context, paging *common.Paging, query *order_item_dto.QueryDTO) ([]model.OrderItem, error) {
	//there will have business logic before getting data list with condition
	data, err := service.orderItemRepo.FindAllWithCondition(ctx, paging, query)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return data, nil
}

func (service *orderItemService) FindOneById(ctx context.Context, id int) (*model.OrderItem, error) {
	//there will have business logic before getting specific data with condition

	data, err := service.orderItemRepo.FindOneWithCondition(ctx, map[string]any{"id": id}, "Order")

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.OrderItemEntity, err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return data, nil

}

func (service *orderItemService) Update(ctx context.Context, id int, dto *order_item_dto.CreateDTO) error {

	if err := dto.Validate(); err != nil {
		return err
	}

	if _, err := service.FindOneById(ctx, id); err != nil {
		return err
	}

	if err := service.orderItemRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *orderItemService) Delete(ctx context.Context, id int) error {

	if _, err := service.FindOneById(ctx, id); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.orderItemRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
