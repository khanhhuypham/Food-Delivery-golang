package middleware

import (
	"Food-Delivery/config"
	"Food-Delivery/entity/model"
	"context"
)

type UserRepository interface {
	FindDataWithCondition(context.Context, map[string]any) (*model.User, error)
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
