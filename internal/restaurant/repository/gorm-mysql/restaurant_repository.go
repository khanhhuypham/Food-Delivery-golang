package gorm_mysql

import (
	restaurant_dto "Food-Delivery/entity/dto/restaurant"
	"Food-Delivery/entity/model"
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
