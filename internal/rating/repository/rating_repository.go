package rating_repository

import (
	rating_dto "Food-Delivery/entity/dto/rating"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ratingRepository struct {
	tableName string
	db        *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *ratingRepository {
	return &ratingRepository{
		tableName: model.Rating{}.TableName(),
		db:        db,
	}
}

func (repo *ratingRepository) FindOneWithCondition(ctx context.Context, condition map[string]any) (*model.Rating, error) {
	db := repo.db.Table(repo.tableName)

	var data model.Rating

	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}

func (repo *ratingRepository) Create(ctx context.Context, dto *rating_dto.CreateDTO) error {
	//apply transaction technique
	db := repo.db.Begin()

	if _, err := repo.FindUserById(ctx, dto.UserId); err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}

	if id := dto.RestaurantId; id != nil {
		if _, err := repo.FindRestaurantById(ctx, *id); err != nil {
			db.Rollback()
			return errors.WithStack(err)
		}
	}

	if id := dto.ItemId; id != nil {
		if _, err := repo.FindItemById(ctx, *id); err != nil {
			db.Rollback()
			return errors.WithStack(err)
		}
	}

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

func (repo *ratingRepository) Update(ctx context.Context, id int, dto *rating_dto.CreateDTO) error {
	//apply transaction technique
	db := repo.db.Begin().Debug()

	updateData := map[string]interface{}{}

	if dto.Like != nil {
		updateData["like"] = dto.Like
	}

	if dto.Score != nil {
		updateData["score"] = *dto.Score
	}
	if dto.Comment != nil {
		updateData["comment"] = *dto.Comment
	}

	if err := db.Table(repo.tableName).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(updateData).Error; err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}

	return nil
}

func (repo *ratingRepository) DeleteLike(ctx context.Context, condition map[string]any) error {

	if err := repo.db.Table(repo.tableName).Where(condition).Delete(&model.Rating{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *ratingRepository) FindRestaurantById(ctx context.Context, id int) (*model.Restaurant, error) {
	var data model.Restaurant

	db := repo.db.Table(data.TableName())

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.RestaurantEntity, err)
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}

func (repo *ratingRepository) FindUserById(ctx context.Context, id int) (*model.User, error) {
	var data model.User

	db := repo.db.Table(data.TableName())

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.UserEntity, err)
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}

func (repo *ratingRepository) FindItemById(ctx context.Context, id int) (*model.Item, error) {
	var data model.Item

	db := repo.db.Table(data.TableName())

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.ItemEntity, err)
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
