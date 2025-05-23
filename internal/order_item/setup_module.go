package order_item_module

import (
	order_item_http_handler "Food-Delivery/internal/order_item/controller/http"
	order_item_repository "Food-Delivery/internal/order_item/repository"
	order_item_service "Food-Delivery/internal/order_item/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, r *gin.RouterGroup) {
	
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
