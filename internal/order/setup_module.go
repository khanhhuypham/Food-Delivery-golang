package order_module

import (
	order_http_handler "Food-Delivery/internal/order/controller/http"
	order_repository "Food-Delivery/internal/order/repository"
	order_service "Food-Delivery/internal/order/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, r *gin.RouterGroup) {

	//dependency of place module
	repo := order_repository.NewOrderRepository(db)
	service := order_service.NewOrderService(repo)
	handler := order_http_handler.NewOrderHandler(service)

	r.POST("/order", handler.Create())
	r.GET("/order", handler.GetAll())
	r.GET("/order/:id", handler.GetOneByID())
	r.GET("/order/change-status/:id", handler.ChangeStatus())

}
