package children_item_repository

import (
	children_item_dto "Food-Delivery/entity/dto/children_item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type childrenItemRepository struct {
	db        *gorm.DB
	tableName string
}

func NewChildrenItemRepository(db *gorm.DB) *childrenItemRepository {

	return &childrenItemRepository{
		db:        db,
		tableName: model.ChildrenItem{}.TableName(),
	}

}

// create place
func (repo *childrenItemRepository) Create(ctx context.Context, dto *children_item_dto.CreateDTO) (*model.ChildrenItem, error) {

	var data model.ChildrenItem

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

func (repo *childrenItemRepository) FindAllWithCondition(
	ctx context.Context,
	query *children_item_dto.QueryDTO,
	keys ...string) ([]model.ChildrenItem, error) {

	var data []model.ChildrenItem
	db := repo.db.Model(&data).Table(repo.tableName).Where("restaurant_id = ?", query.RestaurantId)

	if query.SearchKey != nil {
		db.Where("name LIKE ?", "%"+*query.SearchKey+"%")
	}

	for _, v := range keys {
		db = db.Preload(v)
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}

func (repo *childrenItemRepository) FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.ChildrenItem, error) {
	var data model.ChildrenItem
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
func (repo *childrenItemRepository) DeleteDataWithCondition(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.ChildrenItem{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// update place by condition
func (repo *childrenItemRepository) UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *children_item_dto.CreateDTO) (*model.ChildrenItem, error) {

	var updatedData model.ChildrenItem

	if err := repo.db.Table(repo.tableName).Clauses(clause.Returning{}).Where(condition).Updates(dto).Scan(&updatedData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &updatedData, nil
}
