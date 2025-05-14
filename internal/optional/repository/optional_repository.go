package item_optional_repository

import (
	item_optional_dto "Food-Delivery/entity/dto/item_optional"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type itemOptionalRepository struct {
	db                    *gorm.DB
	tableName             string
	ChildrenItemTableName string
}

func NewItemOptionalRepository(db *gorm.DB) *itemOptionalRepository {

	return &itemOptionalRepository{
		db:                    db,
		tableName:             model.Optional{}.TableName(),
		ChildrenItemTableName: model.ChildrenItem{}.TableName(),
	}

}

// create place
func (repo *itemOptionalRepository) Create(ctx context.Context, dto *item_optional_dto.CreateDTO) (*model.Optional, error) {

	var data model.Optional

	err := copier.Copy(&data, dto)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//apply transaction technique
	db := repo.db.Begin()
	if err := repo.db.Table(repo.tableName).Create(dto).Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	return &data, nil
}

func (repo *itemOptionalRepository) FindAllWithCondition(ctx context.Context, restaurantId int, keys ...string) ([]model.Optional, error) {

	var data []model.Optional
	db := repo.db.Model(&data).Table(repo.tableName).Where("restaurant_id = ?", restaurantId)

	for _, v := range keys {
		db = db.Preload(v)
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}

func (repo *itemOptionalRepository) FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Optional, error) {
	var data model.Optional
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
func (repo *itemOptionalRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.ChildrenItem{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *itemOptionalRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *item_optional_dto.CreateDTO) (*model.Optional, error) {

	var updatedData model.Optional

	if err := repo.db.Table(repo.tableName).Clauses(clause.Returning{}).Where(condition).Updates(dto).Scan(&updatedData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &updatedData, nil
}
