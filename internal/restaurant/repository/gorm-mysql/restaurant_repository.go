package gorm_mysql

import (
	restaurant_dto "Food-Delivery/entity/dto/restaurant"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type restaurantRepository struct {
	tableName string
	db        *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *restaurantRepository {
	restaurant := model.Restaurant{}
	return &restaurantRepository{
		tableName: restaurant.TableName(),
		db:        db,
	}
}

// create place
func (repo *restaurantRepository) Create(ctx context.Context, dto *restaurant_dto.CreateDTO) error {

	//apply transaction technique
	db := repo.db.Begin()
	if err := repo.db.Table(repo.tableName).Create(dto).Error; err != nil {
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

	if err := db.Where("active = ?", true).Count(&data.TotalActive).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.Where("active = ?", false).Count(&data.TotalInActive).Error; err != nil {
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

// Delete place by condition
func (repo *restaurantRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.Restaurant{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *restaurantRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *restaurant_dto.CreateDTO) error {

	if err := repo.db.Table(repo.tableName).Clauses(clause.Returning{}).Where(condition).Updates(dto).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
