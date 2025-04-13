package user_repository

import (
	user_dto "Food-Delivery/entity/dto/user"
	"Food-Delivery/entity/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userRepository struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) *userRepository {
	user := model.User{}
	return &userRepository{db: db, tableName: user.TableName()}
}

func (userRepo *userRepository) Create(ctx context.Context, dto *user_dto.UserCreate) error {
	db := userRepo.db.Begin()

	if err := db.Table(userRepo.tableName).Create(dto).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}

func (userRepo *userRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*model.User, error) {
	var user model.User

	if err := userRepo.db.Table(userRepo.tableName).Where(condition).First(&user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (userRepo *userRepository) DeleteUserWithCondition(ctx context.Context, condition map[string]any) error {

	if err := userRepo.db.Table(userRepo.tableName).Where(condition).Delete(&model.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}
	return nil
}
