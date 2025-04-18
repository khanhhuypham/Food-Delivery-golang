package app_context

import (
	"Food-Delivery/config"
	"Food-Delivery/internal/middleware"
	user_repository "Food-Delivery/internal/user/repository"

	"gorm.io/gorm"
	"log"
)

type MiddleWareProvider interface {
}

type DbContext interface{}

type AppContext interface {
	GetMiddlewareProvider() MiddleWareProvider
	GetDbContext() DbContext
	GetConfig() *config.Config
	GetMsgBroker() MesssagBroker
}

type appContext struct {
	middleWareProvider MiddleWareProvider
	dbContext          DbContext
	config             *config.Config
	messagBroker       MesssagBroker
}

func NewAppContext(db *gorm.DB) AppContext {

	cfg, err := config.LoadConfig("Food-Delivery/config/config-local-yml")
	if err != nil {
		log.Fatalln("db connection err: ", err)
	}

	dbCtx := NewDbContext(db)

	middlewareProvider := middleware.NewMiddlewareManager(cfg, user_repository.NewUserRepository(db))

	natsComp := NewNatsComp()

	return &appContext{
		middleWareProvider: middlewareProvider,
		dbContext:          dbCtx,
		config:             cfg,
		messagBroker:       natsComp,
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

func (ctx *appContext) GetMsgBroker() MesssagBroker {
	return ctx.messagBroker
}
