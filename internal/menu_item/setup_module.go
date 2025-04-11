package menu_item_module

import (
	menu_item_http_handler "Food-Delivery/internal/menu_item/controller/http"
	menu_item_model "Food-Delivery/internal/menu_item/entity/model"

	menu_item_repository "Food-Delivery/internal/menu_item/repository"
	menu_item_service "Food-Delivery/internal/menu_item/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func Setup(db *gorm.DB, r *gin.RouterGroup) {

	// Automatically migrate the schema, this will sync the struct to the DB

	if err := db.AutoMigrate(&menu_item_model.MenuItem{}); err != nil {
		log.Fatalf("could not migrate schema: %v", err)
	}

	//dependency of place module
	repo := menu_item_repository.NewMenuItemRepository(db)
	service := menu_item_service.NewRestaurantService(repo)
	http_handler := menu_item_http_handler.NewRestaurantHandler(service)

	r.POST("/menu-item", http_handler.Create())
	r.GET("/menu-item", http_handler.FindAll())
	r.GET("/menu-item/:id", http_handler.FindOneByID())
	r.PUT("/menu-item/:id", http_handler.Update())
	r.PATCH("/menu-item/:id", http_handler.Update())
	r.DELETE("/menu-item/:id", http_handler.Delete())

}
