package media_repository

import (
	"Food-Delivery/entity/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MediaRepository struct {
	tableName string
	db        *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{
		tableName: model.Media{}.TableName(),
		db:        db,
	}
}

func (repo *MediaRepository) Create(ctx context.Context, data *model.Media) error {

	//apply transaction technique
	db := repo.db.Begin()

	if err := repo.db.Table(repo.tableName).Create(data).Error; err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}

	return nil
}

func (repo *MediaRepository) Delete(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.Restaurant{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil

}

func (repo *MediaRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*model.Media, error) {
	var data model.Media
	db := repo.db.Table(repo.tableName)
	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
