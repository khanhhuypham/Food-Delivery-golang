package order_module

import (
	order_http_handler "Food-Delivery/internal/order/controller/http"
	order_model "Food-Delivery/internal/order/entity/order_model"
	order_repository "Food-Delivery/internal/order/repository"
	order_service "Food-Delivery/internal/order/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func Setup(db *gorm.DB, r *gin.RouterGroup) {
	if err := db.AutoMigrate(&order_model.Order{}); err != nil {
		log.Fatalf("could not migrate schema: %v", err)
	}
	//dependency of place module
	repo := order_repository.NewOrderRepository(db)
	service := order_service.NewOrderService(repo)
	handler := order_http_handler.NewOrderHandler(service)

	r.POST("/order", handler.Create())
	r.GET("/order", handler.GetAll())
	r.GET("/order/:id", handler.GetOneByID())
	r.GET("/order/change-status/:id", handler.ChangeStatus())

}
