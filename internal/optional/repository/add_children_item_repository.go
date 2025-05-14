package item_optional_repository

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
)

func (repo *itemOptionalRepository) FindChildrenItemByIds(ctx context.Context, ids []int) ([]model.ChildrenItem, error) {
	if len(ids) == 0 {
		return []model.ChildrenItem{}, nil
	}

	var data []model.ChildrenItem

	// Use the correct column name
	db := repo.db.Model(&data).Table(repo.ChildrenItemTableName).Where("id IN (?)", ids)

	if err := db.Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}

func (repo *itemOptionalRepository) AddChildrenItemToOptional(ctx context.Context, optionalId int, childrenItemIds []int) (*model.Optional, error) {
	var updatedData model.Optional

	// 📝 Begin a transaction to ensure atomicity
	db := repo.db.Begin()

	// 🔄 Bulk update all children items to associate with optionalId
	if err := db.Table(repo.ChildrenItemTableName).Where("id IN ?", childrenItemIds).Update("optional_id", optionalId).Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	// 🔍 Retrieve the updated Optional data
	if err := db.Table(repo.tableName).Where("id = ?", optionalId).First(&updatedData).Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	// ✅ Commit the transaction if all is good
	if err := db.Commit().Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &updatedData, nil
}
func (repo *itemOptionalRepository) RemoveChildrenItemFromOptional(ctx context.Context, optional *model.Optional) (*model.Optional, error) {
	// ✅ Return early if there are no children items
	if len(optional.ChildrenItems) == 0 {
		return optional, nil
	}

	// 📝 Begin a transaction to ensure atomicity
	db := repo.db.Begin()

	// ✅ Preallocate the slice for better memory management
	childrenIds := make([]int, len(optional.ChildrenItems))
	for i, item := range optional.ChildrenItems {
		childrenIds[i] = item.Id
	}

	// 🔄 Bulk update all children to remove the association
	if err := db.Table(repo.ChildrenItemTableName).Where("id IN ?", childrenIds).Update("optional_id", nil).Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	// 🔍 Fetch the updated `Optional` (correct table this time)
	var updatedOptional model.Optional
	if err := db.Table(repo.tableName).Where("id = ?", optional.Id).First(&updatedOptional).Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	// ✅ Commit the transaction if all is good
	if err := db.Commit().Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &updatedOptional, nil
}
