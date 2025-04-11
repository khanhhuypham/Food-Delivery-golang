package user_repository

import (
	"Food-Delivery/internal/user/entity/dto"
	usermodel "Food-Delivery/internal/user/entity/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userRepository struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db, tableName: usermodel.User{}.TableName()}
}

func (userRepo *userRepository) Create(ctx context.Context, dto *dto.UserCreate) error {
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

func (userRepo *userRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*usermodel.User, error) {
	var user usermodel.User
	//
	//log.Fatal(userRepo.tableName)

	if err := userRepo.db.Table(userRepo.tableName).Where(condition).First(&user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (userRepo *userRepository) DeleteUserWithCondition(ctx context.Context, condition map[string]any) error {

	if err := userRepo.db.Table(userRepo.tableName).Where(condition).Delete(&usermodel.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}
	return nil
}
