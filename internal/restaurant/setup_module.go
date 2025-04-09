package restaurant_module

import (
	restaurant_http "Food-Delivery/internal/restaurant/delivery/http"
	restaurant_model "Food-Delivery/internal/restaurant/model"
	restaurant_repository "Food-Delivery/internal/restaurant/repository/gorm-mysql"
	restaurant_service "Food-Delivery/internal/restaurant/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func SetupRestaurantModule(db *gorm.DB, r *gin.RouterGroup) {
	if err := db.AutoMigrate(&restaurant_model.Restaurant{}); err != nil {
		log.Fatalf("could not migrate schema: %v", err)
	}
	//dependency of place module
	repo := restaurant_repository.NewRestaurantRepository(db)
	service := restaurant_service.NewRestaurantService(repo)
	handler := restaurant_http.NewRestaurantHandler(service)

	r.POST("/restaurant", handler.Create())
	r.GET("/restaurant", handler.GetAll())
	r.GET("/restaurant/:id", handler.GetOneByID())
	r.PUT("/restaurant/:id", handler.Update())
	r.POST("/restaurant/:id", handler.Update())
	r.DELETE("/restaurant/:id", handler.Delete())

}
