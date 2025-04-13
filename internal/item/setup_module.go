package item_module

import (
	menu_item_http_handler "Food-Delivery/internal/item/controller/http"
	menu_item_repository "Food-Delivery/internal/item/repository"
	menu_item_service "Food-Delivery/internal/item/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, r *gin.RouterGroup) {

	//dependency of place module
	repo := menu_item_repository.NewItemRepository(db)
	service := menu_item_service.NewRestaurantService(repo)
	http_handler := menu_item_http_handler.NewRestaurantHandler(service)

	r.POST("/menu-item", http_handler.Create())
	r.GET("/menu-item", http_handler.FindAll())
	r.GET("/menu-item/:id", http_handler.FindOneByID())
	r.PUT("/menu-item/:id", http_handler.Update())
	r.PATCH("/menu-item/:id", http_handler.Update())
	r.DELETE("/menu-item/:id", http_handler.Delete())

}
