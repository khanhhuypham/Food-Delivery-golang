package restaurant_service

import (
	restaurant_dto "Food-Delivery/entity/dto/restaurant"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type RestaurantRepository interface {
	Create(ctx context.Context, dto *restaurant_dto.CreateDTO) error
	ListDataWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *restaurant_dto.QueryDTO,
		keys ...string) ([]model.Restaurant, error)
	FindDataWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Restaurant, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *restaurant_dto.CreateDTO) error
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

type restaurantService struct {
	restaurantRepo RestaurantRepository
}

func NewRestaurantService(restaurantRepo RestaurantRepository) *restaurantService {
	return &restaurantService{restaurantRepo}
}

func (service *restaurantService) Create(ctx context.Context, dto *restaurant_dto.CreateDTO) error {
	//------perform business operation such as validate data
	if err := dto.Validate(); err != nil {
		return err
	}
	//------
	if err := service.restaurantRepo.Create(ctx, dto); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *restaurantService) FindAll(ctx context.Context, paging *common.Paging, filter *restaurant_dto.QueryDTO) ([]model.Restaurant, error) {
	//there will have business logic before getting data list with condition
	restaurants, err := service.restaurantRepo.ListDataWithCondition(ctx, paging, filter, "Category")

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return restaurants, nil
}

func (service *restaurantService) FindOneById(ctx context.Context, id int) (*model.Restaurant, error) {
	//there will have business logic before getting specific data with condition

	restaurant, err := service.restaurantRepo.FindDataWithCondition(ctx, map[string]any{"id": id}, "Category", "Orders", "MenuItems")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(restaurant.TableName(), err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return restaurant, nil
}

func (service *restaurantService) Update(ctx context.Context, id int, dto *restaurant_dto.CreateDTO) error {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return err
	}
	//check the eixstence of data in database
	if _, err := service.restaurantRepo.FindDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return err
	}

	if err := service.restaurantRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *restaurantService) Delete(ctx context.Context, id int) error {
	//check the eixstence of data in database
	if _, err := service.restaurantRepo.FindDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.restaurantRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
