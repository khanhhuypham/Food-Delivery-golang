package item_repository

import (
	menu_item_dto "Food-Delivery/entity/dto/item"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/pkg/errors"
)

// get category
func (repo *itemRepository) FindAllWithCondition(
	ctx context.Context,
	paging *common.Paging,
	query *menu_item_dto.QueryDTO,
	keys ...string) ([]model.Item, error) {

	var data []model.Item

	db := repo.db.Table(repo.tableName)

	////Để không count những record bị  soft delete ta cần dùng Model
	//db = repo.db.Model(&data)

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

func (repo *itemRepository) FindTheMostPopularItem(ctx context.Context, paging *common.Paging, keys ...string) ([]model.Item, error) {

	var data []model.Item

	db := repo.db.Table(repo.tableName)

	// Count total records (without pagination)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

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

func (repo *itemRepository) FindTheMostRecommendedItem(ctx context.Context, paging *common.Paging, keys ...string) ([]model.Item, error) {

	var data []model.Item

	db := repo.db.Table(repo.tableName)

	// Count total records (without pagination)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

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
