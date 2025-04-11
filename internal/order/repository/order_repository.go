package order_repository

import (
	order_dto "Food-Delivery/entity/dto/order"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type orderRepository struct {
	tableName string
	db        *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		tableName: model.Order{}.TableName(),
		db:        db,
	}
}

// create place
func (repo *orderRepository) Create(ctx context.Context, dto *order_dto.CreateDTO) error {

	//apply transaction technique
	db := repo.db.Begin()

	// Ensure rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

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
func (repo *orderRepository) FindAllWithCondition(
	ctx context.Context,
	paging *common.Paging,
	query *order_dto.QueryDTO,
	keys ...string) ([]model.Order, error) {

	var data []model.Order

	db := repo.db.Model(&model.Order{}).Table(repo.tableName)

	if query.SearchKey != nil {
		db.Where("name LIKE ?", "%"+*query.SearchKey+"%")
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

func (repo *orderRepository) FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Order, error) {
	var data model.Order
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
func (repo *orderRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.Order{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *orderRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *order_dto.UpdateDTO) (*model.Order, error) {
	var data model.Order

	err := repo.db.WithContext(ctx).
		Table(repo.tableName).
		Clauses(clause.Returning{}).
		Where(condition).
		Updates(dto).
		Scan(&data).
		Error

	if err != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// Duplicate entry (unique constraint violated)
			if strings.Contains(err.Error(), "idx_menu_item_name") {
				return nil, common.ErrBadRequest(errors.New("menu item name already exists for this restaurant"))
			}

		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
