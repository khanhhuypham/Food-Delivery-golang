package item_optional_module

import (
	item_optional_http "Food-Delivery/internal/optional/controller/http"
	item_optional_repository "Food-Delivery/internal/optional/repository"
	item_optional_service "Food-Delivery/internal/optional/service"
	"Food-Delivery/pkg/app_context"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {
	db := appCtx.GetDbContext().GetMainConnection()

	//dependency of place module
	repo := item_optional_repository.NewItemOptionalRepository(db)
	service := item_optional_service.NewItemOptionalService(repo)
	http_handler := item_optional_http.NewItemOptionalHandler(service)

	r.GET("/item-optional", http_handler.FindAll())
	r.GET("/item-optional/:id", http_handler.FindOneByID())
	r.POST("/item-optional", http_handler.Create())
	r.PUT("/item-optional/:id", http_handler.Update())
	r.DELETE("/item-optional/:id", http_handler.Delete())

}
