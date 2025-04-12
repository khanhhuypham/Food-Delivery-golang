package order_item_repository

import (
	order_item_dto "Food-Delivery/entity/dto/order-item"
	"Food-Delivery/entity/model"

	"Food-Delivery/pkg/common"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderItemRepository struct {
	tableName string
	db        *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *orderItemRepository {
	orderItem := model.OrderItem{}
	return &orderItemRepository{
		tableName: orderItem.TableName(),
		db:        db,
	}
}

// create place
func (repo *orderItemRepository) Create(ctx context.Context, dto *order_item_dto.CreateDTO) error {
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
func (repo *orderItemRepository) FindAllWithCondition(
	ctx context.Context,
	paging *common.Paging,
	query *order_item_dto.QueryDTO,
	keys ...string) ([]model.OrderItem, error) {

	var data []model.OrderItem

	db := repo.db.Table(repo.tableName)

	////Để không count những record bị  soft delete ta cần dùng Model
	//db = repo.db.Model(&data)

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

func (repo *orderItemRepository) FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.OrderItem, error) {
	var data model.OrderItem
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
func (repo *orderItemRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.OrderItem{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *orderItemRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *order_item_dto.CreateDTO) error {

	if err := repo.db.Table(repo.tableName).Clauses(clause.Returning{}).Where(condition).Updates(dto).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
