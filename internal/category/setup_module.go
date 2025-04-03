package categorymodule

import (
	category_http "Food-Delivery/internal/category/delivery/http"
	categorymodel "Food-Delivery/internal/category/model"
	category_repository "Food-Delivery/internal/category/repository"
	category_service "Food-Delivery/internal/category/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func SetupCategoryModule(db *gorm.DB, r *gin.RouterGroup) {

	// Automatically migrate the schema, this will sync the struct to the DB

	if err := db.AutoMigrate(&categorymodel.Category{}); err != nil {
		log.Fatalf("could not migrate schema: %v", err)
	}

	//dependency of place module
	cateRepo := category_repository.NewCategoryRepository(db)
	cateService := category_service.NewCategoryService(cateRepo)
	placeHandler := category_http.NewCategoryHandler(cateService)

	r.POST("/category", placeHandler.Create())
	r.GET("/category", placeHandler.FindAll())
	r.GET("/category/:id", placeHandler.FindOneByID())
	r.PUT("/category/:id", placeHandler.Update())
	r.POST("/category/:id", placeHandler.Update())
	r.DELETE("/category/:id", placeHandler.Delete())

}
