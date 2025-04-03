package user_service

import (
	"Food-Delivery/config"
	usermodel "Food-Delivery/internal/user/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"context"
	"errors"
)

type UserRepository interface {
	Create(ctx context.Context, dto *usermodel.UserCreate) error
	FindDataWithCondition(ctx context.Context, condition map[string]any) (*usermodel.User, error)
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

func (userService *userService) Signup(ctx context.Context, dto *usermodel.UserCreate) error {
	userRepo := userService.userRepo

	user, _ := userRepo.FindDataWithCondition(ctx, map[string]any{
		"email": dto.Email,
	})
	if user != nil {
		return errors.New("email already exists")
	}

	if err := dto.PrepareCreate(); err != nil {
		return errors.New("email and password invalid")
	}

	if err := userRepo.Create(ctx, dto); err != nil {

		return err
	}

	return nil
}

func (userService *userService) SignIn(ctx context.Context, dto *usermodel.UserCreate) (*utils.Token, error) {

	user, err := userService.userRepo.FindDataWithCondition(ctx, map[string]any{
		"email": dto.Email,
	})

	if err != nil {
		return nil, err
	}

	if err := utils.CheckPasswordHash(user.Password, dto.Password); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJwt(utils.TokePayload{Email: user.Email, Role: user.Role}, userService.cfg)

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
		return common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if err := userRepo.DeleteUserWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return err
	}

	return nil
}
