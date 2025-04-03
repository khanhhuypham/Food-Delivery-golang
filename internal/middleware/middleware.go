package middleware

import (
	"Food-Delivery/config"
	usermodel "Food-Delivery/internal/user/model"
	"context"
)

type UserRepository interface {
	FindDataWithCondition(context.Context, map[string]any) (*usermodel.User, error)
}

type middlewareManager struct {
	configuration *config.Config
	userRepo      UserRepository
}

func NewMiddlewareManager(configuration *config.Config, userRepo UserRepository) *middlewareManager {
	return &middlewareManager{
		configuration: configuration,
		userRepo:      userRepo,
	}
}
