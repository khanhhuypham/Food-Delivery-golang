package user_repository

import (
	usermodel "Food-Delivery/internal/user/model"
	"context"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (userRepo *userRepository) Create(ctx context.Context, dto *usermodel.UserCreate) error {
	db := userRepo.db.Begin()
	table_name := usermodel.User{}.TableName()
	if err := db.Take(table_name).Create(dto).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}

func (userRepo *userRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*usermodel.User, error) {
	var user usermodel.User
	tableName := usermodel.User{}.TableName()
	if err := userRepo.db.Table(tableName).Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (userRepo *userRepository) DeleteUserWithCondition(ctx context.Context, condition map[string]any) error {
	tableName := usermodel.User{}.TableName()
	if err := userRepo.db.Table(tableName).Where(condition).Delete(&usermodel.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}
	return nil
}
