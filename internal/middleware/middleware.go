package middleware

import (
	"Food-Delivery/config"
	usermodel "Food-Delivery/internal/user/entity/model"
	"context"
)

type UserRepository interface {
	FindDataWithCondition(context.Context, map[string]any) (*usermodel.User, error)
}

type MiddlewareManager struct {
	configuration *config.Config
	userRepo      UserRepository
}

func NewMiddlewareManager(configuration *config.Config, userRepo UserRepository) *MiddlewareManager {
	return &MiddlewareManager{
		configuration: configuration,
		userRepo:      userRepo,
	}
}
