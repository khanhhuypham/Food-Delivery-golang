package category_repository

import (
	category_dto "Food-Delivery/entity/dto/category"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type categoryRepository struct {
	db        *gorm.DB
	tableName string
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db:        db,
		tableName: model.Category{}.TableName(),
	}

}

// create place
func (repo *categoryRepository) Create(ctx context.Context, dto *category_dto.CreateDto) error {

	//apply transaction technique
	db := repo.db.Begin()
	if err := repo.db.Table(repo.tableName).Create(dto).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

func (repo *categoryRepository) FindAllWithCondition(
	ctx context.Context,
	paging *common.Paging,
	query *category_dto.QueryDTO,
	keys ...string) ([]model.Category, error) {

	var data []model.Category
	db := repo.db.Table(repo.tableName)
	//Để không count những record bị  soft delete ta cần dùng Model
	db = repo.db.Model(&data)

	if v := query.Status; len(v) > 0 {
		db = db.Where("status = ?", v)
	}

	// Count total records (without pagination)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
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
		return nil, common.ErrDB(err)
	}

	return data, nil
}

func (repo *categoryRepository) FindAllByIds(ctx context.Context, ids []int, keys ...string) ([]model.Category, error) {

	var data []model.Category
	// Start with the correct table and entity
	db := repo.db.Model(&data).Table(repo.tableName)

	// Apply preloading for relationships if provided
	for _, v := range keys {
		db = db.Preload(v)
	}

	// Use correct SQL syntax for "IN" clause
	if err := db.Where("id IN (?)", ids).Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}

func (repo *categoryRepository) FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Category, error) {
	var data model.Category
	db := repo.db.Table(repo.tableName)

	for _, v := range keys {
		db.Preload(v)
	}

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(data.TableName(), err)
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}

// Delete place by condition
func (repo *categoryRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.Category{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *categoryRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *category_dto.CreateDto) error {

	if err := repo.db.Table(repo.tableName).Clauses(clause.Returning{}).Where(condition).Updates(dto).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
