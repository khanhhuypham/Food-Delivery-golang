package gorm_mysql

import (
	restaurant_model "Food-Delivery/internal/restaurant/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type restaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *restaurantRepository {
	return &restaurantRepository{db: db}
}

// create place
func (repo *restaurantRepository) Create(ctx context.Context, dto *restaurant_model.RestaurantCreateDTO) error {
	tableName := restaurant_model.Restaurant{}.TableName()
	//apply transaction technique
	db := repo.db.Begin()
	if err := repo.db.Table(tableName).Create(dto).Error; err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}

	return nil
}

// get category
func (repo *restaurantRepository) ListDataWithCondition(
	ctx context.Context,
	paging *common.Paging,
	query *restaurant_model.QueryDTO,
	keys ...string) ([]restaurant_model.Restaurant, error) {

	tableName := restaurant_model.Restaurant{}.TableName()

	var data []restaurant_model.Restaurant

	db := repo.db.Table(tableName)

	////Để không count những record bị  soft delete ta cần dùng Model
	//db = repo.db.Model(&data)

	// Check if Status pointer is not nil and points to a non-empty string
	if query.Status != nil && *query.Status != "" {
		db = db.Where("status = ?", *query.Status)
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

func (repo *restaurantRepository) FindDataWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*restaurant_model.Restaurant, error) {
	var data restaurant_model.Restaurant
	db := repo.db.Table(data.TableName())

	for _, v := range keys {
		db.Preload(v)
	}

	if err := db.Where(condition).First(&data).Error; err != nil {

		return nil, errors.WithStack(err)
	}
	return &data, nil
}

// Delete place by condition
func (repo *restaurantRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {
	tableName := restaurant_model.Restaurant{}.TableName()
	if err := repo.db.Table(tableName).Where(condition).Delete(&restaurant_model.Restaurant{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *restaurantRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *restaurant_model.RestaurantCreateDTO) error {
	tableName := restaurant_model.Restaurant{}.TableName()
	if err := repo.db.Table(tableName).Clauses(clause.Returning{}).Where(condition).Updates(dto).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
