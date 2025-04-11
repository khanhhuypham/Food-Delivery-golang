package menu_item_service

import (
	"Food-Delivery/internal/menu_item/entity/dto"
	menu_item_model "Food-Delivery/internal/menu_item/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type MenuItemRepository interface {
	Create(ctx context.Context, dto *dto.MenuItemCreateDTO) (*menu_item_model.MenuItem, error)
	FindAllWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *dto.QueryDTO,
		keys ...string) ([]menu_item_model.MenuItem, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*menu_item_model.MenuItem, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *dto.MenuItemCreateDTO) (*menu_item_model.MenuItem, error)
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

type menuItemService struct {
	menuItemRepo MenuItemRepository
}

func NewRestaurantService(menuItemRepo MenuItemRepository) *menuItemService {
	return &menuItemService{menuItemRepo}
}

func (service *menuItemService) Create(ctx context.Context, menuItem *dto.MenuItemCreateDTO) (*menu_item_model.MenuItem, error) {
	//------perform business operation such as validate data
	if err := menuItem.Validate(); err != nil {
		return nil, err
	}

	newItem, err := service.menuItemRepo.Create(ctx, menuItem)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return newItem, nil
}

func (service *menuItemService) FindAll(ctx context.Context, paging *common.Paging, query *dto.QueryDTO) ([]menu_item_model.MenuItem, error) {
	//there will have business logic before getting data list with condition
	items, err := service.menuItemRepo.FindAllWithCondition(ctx, paging, query)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return items, nil
}

func (service *menuItemService) FindOneById(ctx context.Context, id int) (*menu_item_model.MenuItem, error) {
	//there will have business logic before getting specific data with condition

	item, err := service.menuItemRepo.FindOneWithCondition(ctx, map[string]any{"id": id}, "Restaurant")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(menu_item_model.EntityName, err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return item, nil
}

func (service *menuItemService) Update(ctx context.Context, id int, dto *dto.MenuItemCreateDTO) (*menu_item_model.MenuItem, error) {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	//check the eixstence of data in database
	if _, err := service.menuItemRepo.FindOneWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return nil, err
	}

	updatedItem, err := service.menuItemRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return updatedItem, nil
}

func (service *menuItemService) Delete(ctx context.Context, id int) error {
	//check the eixstence of data in database
	if _, err := service.menuItemRepo.FindOneWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.menuItemRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
