package gorm_mysql

import (
	"Food-Delivery/entity/constant"
	restaurant_dto "Food-Delivery/entity/dto/restaurant"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/pkg/errors"
)

// get category
func (repo *restaurantRepository) ListDataWithCondition(
	ctx context.Context,
	paging *common.Paging,
	query *restaurant_dto.QueryDTO,
	keys ...string) ([]model.Restaurant, error) {

	var data []model.Restaurant

	db := repo.db.Table(repo.tableName).Model(&model.Restaurant{})

	// Check if Status pointer is not nil and points to a non-empty string
	if query.Status != nil && *query.Status != "" {
		db = db.Where("status = ?", *query.Status)
	}

	if query.Active != nil {
		db = db.Where("active = ?", *query.Active)
	}

	if query.SearchKey != nil && *query.SearchKey != "" {
		db = db.Debug().Where("name LIKE ?", "%"+*query.SearchKey+"%")
	}

	// Count total records (without pagination)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// Apply cursor-based pagination if available
	for _, v := range keys {
		db = db.Preload(v)
	}

	// Apply offset and limit for pagination
	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset).Limit(paging.Limit)

	// Fetch the data
	if err := db.Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}

// get category
func (repo *restaurantRepository) GetStatistic() (*restaurant_dto.Statistic, error) {

	var data restaurant_dto.Statistic

	db := repo.db.Table(repo.tableName)

	////Để không count những record bị  soft delete ta cần dùng Model
	//db = repo.db.Model(&data)

	// Count total records (without pagination)
	if err := db.Count(&data.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	//RESTAURANT_STATUS_OPEN                 RestaurantStatus = 1 //- The store is currently operating and accepting orders.
	//RESTAURANT_STATUS_CLOSED               RestaurantStatus = 2 // - The store is not operating (e.g., outside business hours)
	//RESTAURANT_STATUS_TEMPORARILY_CLOSED   RestaurantStatus = 3 //– Closed due to temporary reasons (e.g., holiday, maintenance).
	//RESTAURANT_STATUS_LIMITED_AVAILABILITY RestaurantStatus = 4 // -Temporarily not accepting orders due to high load
	//RESTAURANT_STATUS_SUSPENDED

	if err := db.Where("status = ?", constant.RESTAURANT_STATUS_OPEN).Count(&data.TotalActive).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.Where("status = ?", constant.RESTAURANT_STATUS_OPEN).Count(&data.TotalInActive).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}

func (repo *restaurantRepository) FindDataWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Restaurant, error) {
	var data model.Restaurant
	db := repo.db.Table(repo.tableName)

	for _, v := range keys {
		db.Preload(v)
	}

	if err := db.Where(condition).First(&data).Error; err != nil {

		return nil, errors.WithStack(err)
	}
	return &data, nil
}

func (repo *restaurantRepository) FindTheMostPopularRestaurant(ctx context.Context, paging *common.Paging, keys ...string) ([]model.Restaurant, error) {

	var data []model.Restaurant

	db := repo.db.Table(repo.tableName)

	// Count total records (without pagination)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	for _, v := range keys {
		db = db.Preload(v)
	}

	// Apply offset and limit for pagination
	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset).Limit(paging.Limit)

	// Fetch the data
	if err := db.Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}

func (repo *restaurantRepository) FindTheMostRecommendedRestaurant(ctx context.Context, paging *common.Paging, keys ...string) ([]model.Restaurant, error) {

	var data []model.Restaurant

	db := repo.db.Table(repo.tableName)

	// Count total records (without pagination)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	for _, v := range keys {
		db = db.Preload(v)
	}

	// Apply offset and limit for pagination
	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset).Limit(paging.Limit)

	// Fetch the data
	if err := db.Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}
