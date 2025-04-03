package category_repository

import (
	categorymodel "Food-Delivery/internal/category/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

// create place
func (repo *categoryRepository) Create(ctx context.Context, dto *categorymodel.CategoryCreateDto) error {
	tableName := categorymodel.Category{}.TableName()
	//apply transaction technique
	db := repo.db.Begin()
	if err := repo.db.Table(tableName).Create(dto).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

// get category
func (repo *categoryRepository) ListDataWithCondition(
	ctx context.Context,
	paging *common.Paging,
	query *categorymodel.QueryDTO,
	keys ...string) ([]categorymodel.Category, error) {

	var data []categorymodel.Category
	db := repo.db.Table(categorymodel.Category{}.TableName())
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

func (repo *categoryRepository) FindDataWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*categorymodel.Category, error) {
	var data categorymodel.Category
	db := repo.db.Table(data.TableName())

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
	tableName := categorymodel.Category{}.TableName()
	if err := repo.db.Table(tableName).Where(condition).Delete(&categorymodel.Category{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *categoryRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *categorymodel.CategoryCreateDto) error {
	tableName := categorymodel.Category{}.TableName()
	if err := repo.db.Table(tableName).Clauses(clause.Returning{}).Where(condition).Updates(dto).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
