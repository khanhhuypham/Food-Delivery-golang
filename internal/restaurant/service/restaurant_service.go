package restaurant_service

import (
	categorymodel "Food-Delivery/internal/category/model"
	restaurant_model "Food-Delivery/internal/restaurant/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/google/uuid"
)

type RestaurantRepository interface {
	Create(ctx context.Context, dto *restaurant_model.RestaurantCreateDTO) error
	ListDataWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *restaurant_model.QueryDTO,
		keys ...string) ([]restaurant_model.Restaurant, error)
	FindDataWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*restaurant_model.Restaurant, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *restaurant_model.RestaurantCreateDTO) error
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

type restaurantService struct {
	restaurantRepo RestaurantRepository
}

func NewRestaurantService(restaurantRepo RestaurantRepository) *restaurantService {
	return &restaurantService{restaurantRepo}
}

func (service *restaurantService) Create(ctx context.Context, cate *restaurant_model.RestaurantCreateDTO) error {
	//------perform business operation such as validate data
	if err := cate.Validate(); err != nil {
		return common.ErrBadRequest(err)
	}
	//------
	if err := service.restaurantRepo.Create(ctx, cate); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *restaurantService) FindAll(ctx context.Context, paging *common.Paging, filter *restaurant_model.QueryDTO) ([]restaurant_model.Restaurant, error) {
	//there will have business logic before getting data list with condition
	restaurants, err := service.restaurantRepo.ListDataWithCondition(ctx, paging, filter, "category")

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return restaurants, nil
}

func (service *restaurantService) FindOneById(ctx context.Context, id uuid.UUID) (*restaurant_model.Restaurant, error) {
	//there will have business logic before getting specific data with condition

	restaurant, err := service.restaurantRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (service *restaurantService) Update(ctx context.Context, id uuid.UUID, dto *restaurant_model.RestaurantCreateDTO) error {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return err
	}
	//check the eixstence of data in database
	_, err := service.restaurantRepo.FindDataWithCondition(ctx, map[string]any{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(categorymodel.EntityName, err)
	}

	if err := service.restaurantRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *restaurantService) Delete(ctx context.Context, id uuid.UUID) error {
	//check the existence of data in database
	_, err := service.restaurantRepo.FindDataWithCondition(ctx, map[string]any{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(categorymodel.EntityName, err)
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.restaurantRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
