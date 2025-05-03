package order_service

import (
	order_dto "Food-Delivery/entity/dto/order"
	order_item_dto "Food-Delivery/entity/dto/order-item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, dto *order_dto.CreateDTO) error
	FindAllWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *order_dto.QueryDTO,
		keys ...string) ([]model.Order, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Order, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *order_dto.UpdateDTO) (*model.Order, error)
	//AddItemsToOrder(ctx context.Context, orderId int, items []order_item_dto.CreateDTO) (*model.Order, error)
}

type OrderItemService interface {
	Create(ctx context.Context, dto *order_item_dto.CreateDTO) error
	Update(ctx context.Context, id int, dto *order_item_dto.CreateDTO) error
	Delete(ctx context.Context, id int) error
}

type orderService struct {
	orderRepo OrderRepository
}

func NewOrderService(orderRepo OrderRepository) *orderService {
	return &orderService{orderRepo}
}

func (service *orderService) Create(ctx context.Context, data *order_dto.CreateDTO) error {
	//------perform business operation such as validate data
	if err := data.Validate(); err != nil {
		return err
	}

	err := service.orderRepo.Create(ctx, data)

	if err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}

	var items []order_item_dto.CreateDTO

	for _, item := range data.Items {
		items = append(items, order_item_dto.CreateDTO{
			OrderId:  1,
			ItemId:   item.ItemId,
			Quantity: item.Quantity,
		})
	}

	//order, err = service.AddItemToOrder(ctx, order.Id, items)

	if err != nil {
		return err
	}

	return nil
}

func (service *orderService) FindAll(ctx context.Context, paging *common.Paging, query *order_dto.QueryDTO) ([]model.Order, error) {
	//there will have business logic before getting data list with condition
	items, err := service.orderRepo.FindAllWithCondition(ctx, paging, query)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return items, nil
}

func (service *orderService) FindOneById(ctx context.Context, id int) (*model.Order, error) {
	//there will have business logic before getting specific data with condition

	data, err := service.orderRepo.FindOneWithCondition(ctx, map[string]any{"id": id}, "Restaurant", "User")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.ItemEntity, err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return data, nil
}

func (service *orderService) ChangeStatus(ctx context.Context, id int, dto *order_dto.UpdateDTO) (*model.Order, error) {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	//check the eixstence of data in database
	if _, err := service.orderRepo.FindOneWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return nil, err
	}

	updatedItem, err := service.orderRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return updatedItem, nil
}
