package app_context

import (
	"Food-Delivery/config"
	"Food-Delivery/internal/middleware"
	user_repository "Food-Delivery/internal/user/repository"
	"gorm.io/gorm"
)

type MiddleWareProvider interface {
}

type DbContext interface {
	GetMainConnection() *gorm.DB
}

type AppContext interface {
	GetMiddlewareProvider() MiddleWareProvider
	GetDbContext() DbContext
	GetConfig() *config.Config
	GetMsgBroker() MesssageBroker
}

type appContext struct {
	middleWareProvider MiddleWareProvider
	dbContext          DbContext
	config             *config.Config
	messageBroker      MesssageBroker
}

func NewAppContext(cfg *config.Config, db *gorm.DB) AppContext {

	dbCtx := NewDbContext(db)

	middlewareProvider := middleware.NewMiddlewareManager(cfg, user_repository.NewUserRepository(db))

	natsComp := NewNatsComp()

	return &appContext{
		middleWareProvider: middlewareProvider,
		dbContext:          dbCtx,
		config:             cfg,
		messageBroker:      natsComp,
	}

}

func (ctx *appContext) GetMiddlewareProvider() MiddleWareProvider {
	return ctx.middleWareProvider
}

func (ctx *appContext) GetDbContext() DbContext {
	return ctx.dbContext
}

func (ctx *appContext) GetConfig() *config.Config {
	return ctx.config
}

func (ctx *appContext) GetMsgBroker() MesssageBroker {
	return ctx.messageBroker
}
