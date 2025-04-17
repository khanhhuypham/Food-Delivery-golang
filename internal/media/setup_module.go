package upload_module

import (
	"Food-Delivery/config"
	media_http "Food-Delivery/internal/media/controller/http"
	media_repository "Food-Delivery/internal/media/repository"
	media_service "Food-Delivery/internal/media/service"
	"Food-Delivery/pkg/upload"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, r *gin.RouterGroup, cfg *config.Config) {

	//Declare s3
	s3Provider := upload.NewS3Provider(cfg)

	//Declare service
	repo := media_repository.NewMediaRepository(db)
	service := media_service.NewMediaService(repo, s3Provider)
	handler := media_http.NewMediaHandler(s3Provider, service)

	r.POST("/upload", handler.Upload())
}
