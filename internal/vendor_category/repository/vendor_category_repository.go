package vendor_category_repository

import (
	vendor_category_dto "Food-Delivery/entity/dto/vendor_category"
	"Food-Delivery/entity/model"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type vendorCategoryRepository struct {
	tableName string
	db        *gorm.DB
}

func NewVendorCategoryRepository(db *gorm.DB) *vendorCategoryRepository {
	return &vendorCategoryRepository{
		tableName: model.VendorCategory{}.TableName(),
		db:        db,
	}
}

func (repo *vendorCategoryRepository) FindAllByRestaurantId(ctx context.Context, restaurantId int, keys ...string) ([]model.VendorCategory, error) {

	var data []model.VendorCategory

	db := repo.db.Model(&model.VendorCategory{}).Table(repo.tableName)

	for _, v := range keys {
		db = db.Preload(v)
	}

	// Fetch the data
	if err := db.Where("restaurant_id = ?", restaurantId).Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}

func (repo *vendorCategoryRepository) FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.VendorCategory, error) {
	db := repo.db.Table(repo.tableName)

	var data model.VendorCategory

	for _, v := range keys {
		db = db.Preload(v)
	}

	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}

func (repo *vendorCategoryRepository) Create(ctx context.Context, dto *vendor_category_dto.CreateDTO) (*model.VendorCategory, error) {
	var data model.VendorCategory
	copier.Copy(&data, dto)

	db := repo.db.Begin()
	if err := repo.db.Create(&data).Error; err != nil {
		db.Rollback()
		return nil, errors.WithStack(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, errors.WithStack(err)
	}

	return &data, nil
}

// update place by condition
func (repo *vendorCategoryRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *vendor_category_dto.UpdateDTO) (*model.VendorCategory, error) {
	var updatedData model.VendorCategory

	if err := repo.db.Table(repo.tableName).Clauses(clause.Returning{}).Where(condition).Updates(dto).Scan(&updatedData).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &updatedData, nil
}

// Delete place by condition
func (repo *vendorCategoryRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.Restaurant{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
