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
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *restaurant_dto.CreateDTO) error
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error

	ListDataWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *restaurant_dto.QueryDTO,
		keys ...string) ([]model.Restaurant, error)
	GetStatistic() (*restaurant_dto.Statistic, error)
	FindDataWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Restaurant, error)
	FindTheMostPopularRestaurant(ctx context.Context, paging *common.Paging, keys ...string) ([]model.Restaurant, error)
	FindTheMostRecommendedRestaurant(ctx context.Context, paging *common.Paging, keys ...string) ([]model.Restaurant, error)
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

func (service *restaurantService) FindAll(ctx context.Context, paging *common.Paging, filter *restaurant_dto.QueryDTO) ([]model.Restaurant, *restaurant_dto.Statistic, error) {

	if filter.Status != nil && !filter.Status.IsValid() {
		return nil, nil, common.ErrBadRequest(errors.New("status is invalid"))
	}

	//there will have business logic before getting data list with condition
	restaurants, err := service.restaurantRepo.ListDataWithCondition(ctx, paging, filter, "Rating")

	if err != nil {
		return nil, nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	statistic, err := service.restaurantRepo.GetStatistic()

	if err != nil {
		return nil, nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return restaurants, statistic, nil
}

func (service *restaurantService) FindOneById(ctx context.Context, id int) (*model.Restaurant, error) {
	//there will have business logic before getting specific data with condition

	restaurant, err := service.restaurantRepo.FindDataWithCondition(ctx, map[string]any{"id": id}, "VendorCategory")
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

func (service *restaurantService) FindTheMostPopularRestaurant(ctx context.Context, paging *common.Paging) ([]model.Restaurant, error) {
	//there will have business logic before getting data list with condition
	items, err := service.restaurantRepo.FindTheMostPopularRestaurant(ctx, paging, "Rating")

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return items, nil
}

func (service *restaurantService) FindTheMostRecommendedRestaurant(ctx context.Context, paging *common.Paging) ([]model.Restaurant, error) {
	//there will have business logic before getting data list with condition
	items, err := service.restaurantRepo.FindTheMostRecommendedRestaurant(ctx, paging, "Rating")

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return items, nil
}
