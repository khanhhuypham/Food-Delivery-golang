package restaurant_module

import (
	"Food-Delivery/internal/restaurant/delivery/grpc-client"
	restaurant_http "Food-Delivery/internal/restaurant/delivery/http"
	"Food-Delivery/pkg/app_context"

	restaurant_repository "Food-Delivery/internal/restaurant/repository/gorm-mysql"
	restaurant_service "Food-Delivery/internal/restaurant/service"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {
	db := appCtx.GetDbContext().GetMainConnection()
	cfg := appCtx.GetConfig()
	//dependency of place module
	repo := restaurant_repository.NewRestaurantRepository(db)
	service := restaurant_service.NewRestaurantService(repo)
	// categoryRPCClient := rpcclient.NewCategoryRPCClient(appCtx.GetConfig().CategoryServiceURL)
	categoryGRPCClient := category_grpc_client.NewCategoryGRPCClient(cfg.Grpc.Url)
	handler := restaurant_http.NewRestaurantHandler(service, categoryGRPCClient)
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

	r.POST("/restaurant", handler.Create())
	r.GET("/restaurant", handler.GetAll())
	r.GET("/restaurant/the-most-popular", handler.FindTheMostPopularRestaurant())
	r.GET("/restaurant/the-most-recommended", handler.FindTheMostRecommendedRestaurant())
	r.GET("/restaurant/:id", handler.GetOneByID())
	r.PUT("/restaurant/:id", handler.Update())
	r.PATCH("/restaurant/:id", handler.Update())
	r.DELETE("/restaurant/:id", handler.Delete())

}
