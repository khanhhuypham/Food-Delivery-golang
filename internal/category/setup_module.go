package category_module

import (
	"Food-Delivery/entity/model"
	category_http "Food-Delivery/internal/category/controller/http"
	rpc_category_handler "Food-Delivery/internal/category/controller/rpc"
	category_repository "Food-Delivery/internal/category/repository"
	category_service "Food-Delivery/internal/category/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func Setup(db *gorm.DB, r *gin.RouterGroup) {

	// Automatically migrate the schema, this will sync the struct to the DB

	if err := db.AutoMigrate(&model.Category{}); err != nil {
		log.Fatalf("could not migrate schema: %v", err)
	}

	//dependency of place module
	cateRepo := category_repository.NewCategoryRepository(db)
	cateService := category_service.NewCategoryService(cateRepo)
	http_handler := category_http.NewCategoryHandler(cateService)
	rpc_handler := rpc_category_handler.NewRPCCategoryHandler(cateService)

	r.POST("/category", http_handler.Create())
	r.GET("/category", http_handler.FindAll())
	r.GET("/category/:id", http_handler.FindOneByID())
	r.PUT("/category/:id", http_handler.Update())
	r.POST("/category/:id", http_handler.Update())
	r.DELETE("/category/:id", http_handler.Delete())
	r.POST("/rpc/categories/find-by-ids", rpc_handler.GetByIds())

}
