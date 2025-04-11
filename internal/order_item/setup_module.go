package order_item_module

import (
	order_item_http_handler "Food-Delivery/internal/order_item/controller/http"
	"Food-Delivery/internal/order_item/entity/order_item_model"
	order_item_repository "Food-Delivery/internal/order_item/repository"
	order_item_service "Food-Delivery/internal/order_item/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func Setup(db *gorm.DB, r *gin.RouterGroup) {
	if err := db.AutoMigrate(&order_item_model.OrderItem{}); err != nil {
		log.Fatalf("could not migrate schema: %v", err)
	}
	////dependency of place module
	repo := order_item_repository.NewOrderItemRepository(db)
	service := order_item_service.NewOrderItemService(repo)
	handler := order_item_http_handler.NewOrderItemHandler(service)

	r.POST("/order-item", handler.Create())
	r.GET("/order-item", handler.GetAll())
	r.GET("/order-item/:id", handler.GetOneByID())
	r.PUT("/order-item/:id", handler.Update())
	r.POST("/order-item/:id", handler.Update())
	r.DELETE("/order-item/:id", handler.Delete())

}
