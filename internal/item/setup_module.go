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

	r.POST("/item", http_handler.Create())
	r.PUT("/item/:id", http_handler.Update())
	r.PATCH("/item/:id", http_handler.Update())
	r.DELETE("/item/:id", http_handler.Delete())

	/*
	  - Short by:
	          + Polular: ->  Bảng nào?
	          + Free delivery  -> ??
	          + Nearest me -> ?
	          + Cost low to high: -> ??
	          + Delivery time: -> ??
	      - Rating: AVG(restaurant_ratings.point)?
	      - Price Range: foods.price?
	*/

	r.GET("/item", http_handler.FindAll())
	r.GET("/item/:id", http_handler.FindOneByID())
	r.GET("/item/the-most-popular", http_handler.FindTheMostPopularItem())
	r.GET("/item/the-most-recommended", http_handler.FindTheMostRecommendedItem())

}
