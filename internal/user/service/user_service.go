package user_service

import (
	"Food-Delivery/config"
	"Food-Delivery/entity/constant"
	user_dto "Food-Delivery/entity/dto/user"
	"Food-Delivery/entity/model"

	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, dto *user_dto.UserCreate) error
	FindDataWithCondition(ctx context.Context, condition map[string]any) (*model.User, error)
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

func (service *userService) Signup(ctx context.Context, dto *user_dto.UserCreate) error {

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
		if existUser.Status == constant.USER_DELETED || existUser.Status == constant.USER_BANNED {
			return common.ErrBadRequest(constant.ErrUserBannedOrDeleted)
		}
		return common.ErrBadRequest(constant.ErrEmailAlreadyExists).WithDebug(err.Error())
	}

	if err := dto.PrepareCreate(); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}

	if err := service.userRepo.Create(ctx, dto); err != nil {
		return err
	}

	return nil
}

func (service *userService) SignIn(ctx context.Context, dto *user_dto.UserLogin) (*utils.Token, error) {

	if err := dto.Validate(); err != nil {
		return nil, common.ErrBadRequest(err)
	}

	user, err := service.userRepo.FindDataWithCondition(ctx, map[string]any{
		"email": dto.Email,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.UserEntity, err)
		}
		return nil, common.ErrInternal(err)
	}

	if user.Status == constant.USER_DELETED || user.Status == constant.USER_BANNED {
		return nil, common.ErrBadRequest(constant.ErrUserBannedOrDeleted)
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
		return common.ErrEntityNotFound(model.UserEntity, err)
	}

	if err := userRepo.DeleteUserWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return err
	}

	return nil
}

func (userService *userService) FindById(ctx context.Context, id int) (*model.User, error) {
	userRepo := userService.userRepo

	user, err := userRepo.FindDataWithCondition(ctx, map[string]any{
		"id": id,
	})
	if err != nil {
		return nil, common.ErrEntityNotFound(model.UserEntity, err)
	}

	return user, nil
}
