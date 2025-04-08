package user_module

import (
	"Food-Delivery/config"
	"Food-Delivery/internal/middleware"
	user_http "Food-Delivery/internal/user/delivery/http"
	usermodel "Food-Delivery/internal/user/model"
	user_repository "Food-Delivery/internal/user/repository"
	user_service "Food-Delivery/internal/user/service"
	"Food-Delivery/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func SetupUserModule(db *gorm.DB, r *gin.RouterGroup, cfg *config.Config, hasher *utils.Hasher, middleWare *middleware.MiddlewareManager) {
	if err := db.AutoMigrate(&usermodel.User{}); err != nil {
		log.Fatalf("could not migrate schema: %v", err)
	}
	//dependency of place module
	repo := user_repository.NewUserRepository(db)
	service := user_service.NewUserService(repo, cfg)
	handler := user_http.NewUserHandler(service, hasher)

	r.POST("/auth/sign-up", handler.Signup())
	r.POST("/auth/sign-in", handler.SignIn())
	r.GET("/profile", middleWare.RequireAuth(), handler.GetProfileAPI())

}
