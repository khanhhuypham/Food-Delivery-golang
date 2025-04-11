package user_service

import (
	"Food-Delivery/config"

	"Food-Delivery/internal/user/entity/dto"
	"Food-Delivery/internal/user/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, dto *dto.UserCreate) error
	FindDataWithCondition(ctx context.Context, condition map[string]any) (*user_model.User, error)
	DeleteUserWithCondition(ctx context.Context, condition map[string]any) error
}

type userService struct {
	cfg      *config.Config
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository, cfg *config.Config) *userService {
	return &userService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (service *userService) Signup(ctx context.Context, dto *dto.UserCreate) error {

	if err := dto.Validate(); err != nil {
		return common.ErrBadRequest(err).WithDebug(err.Error())
	}

	existUser, err := service.userRepo.FindDataWithCondition(ctx, map[string]any{
		"email": dto.Email,
	})

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.ErrInternal(err).WithDebug(err.Error())
	}

	if existUser != nil {
		if existUser.Status == user_model.StatusDeleted || existUser.Status == user_model.StatusBanned {
			return common.ErrBadRequest(user_model.ErrUserBannedOrDeleted)
		}
		return common.ErrBadRequest(user_model.ErrEmailAlreadyExists).WithDebug(err.Error())
	}

	if err := dto.PrepareCreate(); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}

	if err := service.userRepo.Create(ctx, dto); err != nil {
		return err
	}

	return nil
}

func (service *userService) SignIn(ctx context.Context, dto *dto.UserLogin) (*utils.Token, error) {

	if err := dto.Validate(); err != nil {
		return nil, common.ErrBadRequest(err)
	}

	user, err := service.userRepo.FindDataWithCondition(ctx, map[string]any{
		"email": dto.Email,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(user_model.EntityName, err)
		}
		return nil, common.ErrInternal(err)
	}

	if user.Status == user_model.StatusDeleted || user.Status == user_model.StatusBanned {
		return nil, common.ErrBadRequest(user_model.ErrUserBannedOrDeleted)
	}

	if err := utils.CheckPasswordHash(dto.Password, user.Password); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJwt(utils.TokePayload{Email: user.Email, Role: string(user.Role)}, service.cfg)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return token, nil
}

func (userService *userService) DeleteUserById(ctx context.Context, id int) error {
	userRepo := userService.userRepo
	_, err := userRepo.FindDataWithCondition(ctx, map[string]any{
		"id": id,
	})

	if err != nil {
		return common.ErrEntityNotFound(user_model.EntityName, err)
	}

	if err := userRepo.DeleteUserWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return err
	}

	return nil
}

func (userService *userService) FindById(ctx context.Context, id int) (*user_model.User, error) {
	userRepo := userService.userRepo

	user, err := userRepo.FindDataWithCondition(ctx, map[string]any{
		"id": id,
	})
	if err != nil {
		return nil, common.ErrEntityNotFound(user_model.EntityName, err)
	}

	return user, nil
}
