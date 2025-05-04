package item_repository

import (
	menu_item_dto "Food-Delivery/entity/dto/item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type itemRepository struct {
	tableName string
	db        *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {

	item := model.Item{}
	return &itemRepository{
		tableName: item.TableName(),
		db:        db,
	}
}

// create place
func (repo *itemRepository) Create(ctx context.Context, dto *menu_item_dto.CreateDTO) (*model.Item, error) {
	var newItem model.Item

	// Start the transaction
	db := repo.db.Begin()

	// Ensure rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	// Attempt to create and scan the new menu item
	err := repo.db.Table(repo.tableName).Create(dto).Scan(&newItem).Error

	// Handle errors
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// Duplicate entry error handling
			if strings.Contains(err.Error(), "idx_menu_item_name") {
				db.Rollback()
				return nil, common.ErrBadRequest(errors.New("menu item name already exists for this restaurant"))
			}
		}
		// Rollback transaction and return error
		db.Rollback()
		return nil, errors.WithStack(err)
	}

	// Commit the transaction
	if err := db.Commit().Error; err != nil {
		// If commit fails, ensure rollback
		db.Rollback()
		return nil, errors.WithStack(err)
	}

	// Return the created menu item
	return &newItem, nil
}

// create place
func (repo *itemRepository) BatchCreate(ctx context.Context, dtos []menu_item_dto.CreateDTO) error {
	var data []model.Item

	for _, dto := range dtos {
		var item model.Item
		if err := copier.Copy(&item, dto); err != nil {
			return errors.WithStack(err)
		}
		data = append(data, item)
	}

	db := repo.db.WithContext(ctx).Begin()
	if err := db.CreateInBatches(&data, 100).Error; err != nil { // 100 is an example batch size
		db.Rollback()
		return errors.WithStack(err)
	}

	if err := db.Commit().Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Delete place by condition
func (repo *itemRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.Item{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *itemRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *menu_item_dto.UpdateDTO) (*model.Item, error) {
	var updatedData model.Item

	err := repo.db.WithContext(ctx).
		Table(repo.tableName).
		Clauses(clause.Returning{}).
		Where(condition).
		Updates(dto).
		Scan(&updatedData).
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
	return &updatedData, nil
}
