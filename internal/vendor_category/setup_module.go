package vendor_category_module

import (
	vendor_category_http_handler "Food-Delivery/internal/vendor_category/controller/http"

	vendor_category_repository "Food-Delivery/internal/vendor_category/repository"
	vendor_category_service "Food-Delivery/internal/vendor_category/service"
	"Food-Delivery/pkg/app_context"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {
	db := appCtx.GetDbContext().GetMainConnection()
	_ = appCtx.GetConfig()

	repo := vendor_category_repository.NewVendorCategoryRepository(db)
	service := vendor_category_service.NewVendorCategoryService(repo)
	http_handler := vendor_category_http_handler.NewVendorCategoryHandler(service)

	r.GET("/vendor-category", http_handler.FindAll())
	r.GET("/vendor-category/:id", http_handler.FindOneByID())
	r.POST("/vendor-category", http_handler.Create())
	r.PUT("/vendor-category/:id", http_handler.Update())
	r.POST("/vendor-category/:id", http_handler.Update())
	r.DELETE("/vendor-category/:id", http_handler.Delete())

}
