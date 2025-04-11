package order_service

import (
	menu_item_model "Food-Delivery/internal/menu_item/entity/model"
	"Food-Delivery/internal/order/entity/dto"
	order_model "Food-Delivery/internal/order/entity/order_model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, dto *dto.OrderCreateDTO) error
	FindAllWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *dto.QueryDTO,
		keys ...string) ([]order_model.Order, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*order_model.Order, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *dto.OrderUpdateDTO) (*order_model.Order, error)
}

type orderService struct {
	orderRepo OrderRepository
}

func NewOrderService(orderRepo OrderRepository) *orderService {
	return &orderService{orderRepo}
}

func (service *orderService) Create(ctx context.Context, data *dto.OrderCreateDTO) error {
	//------perform business operation such as validate data
	if err := data.Validate(); err != nil {
		return err
	}

	if err := service.orderRepo.Create(ctx, data); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *orderService) FindAll(ctx context.Context, paging *common.Paging, query *dto.QueryDTO) ([]order_model.Order, error) {
	//there will have business logic before getting data list with condition
	items, err := service.orderRepo.FindAllWithCondition(ctx, paging, query)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return items, nil
}

func (service *orderService) FindOneById(ctx context.Context, id int) (*order_model.Order, error) {
	//there will have business logic before getting specific data with condition

	data, err := service.orderRepo.FindOneWithCondition(ctx, map[string]any{"id": id}, "Restaurant", "User")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(menu_item_model.EntityName, err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return data, nil
}

func (service *orderService) ChangeStatus(ctx context.Context, id int, dto *dto.OrderUpdateDTO) (*order_model.Order, error) {
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
