package category_module

import (
	category_http "Food-Delivery/internal/category/controller/http"
	rpc_category_handler "Food-Delivery/internal/category/controller/rpc"
	category_repository "Food-Delivery/internal/category/repository"
	category_service "Food-Delivery/internal/category/service"
	media_repository "Food-Delivery/internal/media/repository"
	media_service "Food-Delivery/internal/media/service"
	"Food-Delivery/pkg/app_context"
	"Food-Delivery/pkg/upload"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {
	db := appCtx.GetDbContext().GetMainConnection()
	//Declare s3
	s3Provider := upload.NewS3Provider(appCtx.GetConfig())
	repo := media_repository.NewMediaRepository(db)
	mediaService := media_service.NewMediaService(repo, s3Provider)

	//dependency of place module
	cateRepo := category_repository.NewCategoryRepository(db)
	cateService := category_service.NewCategoryService(cateRepo, mediaService)
	http_handler := category_http.NewCategoryHandler(cateService)
	rpc_handler := rpc_category_handler.NewRPCCategoryHandler(cateService)

	r.POST("/category", http_handler.Create())
	r.GET("/category", http_handler.FindAll())
	r.GET("/category/:id", http_handler.FindOneByID())
	r.PUT("/category/:id", http_handler.Update())
	r.PATCH("/category/:id", http_handler.Update())
	r.DELETE("/category/:id", http_handler.Delete())
	r.POST("/rpc/categories/find-by-ids", rpc_handler.GetByIds())

}
