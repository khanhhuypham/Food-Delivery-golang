package children_item_module

import (
	children_item_http "Food-Delivery/internal/children_item/controller/http"
	children_item_repository "Food-Delivery/internal/children_item/repository"
	children_item_service "Food-Delivery/internal/children_item/service"
	"Food-Delivery/pkg/app_context"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {
	db := appCtx.GetDbContext().GetMainConnection()
	_ = appCtx.GetConfig()

	//dependency of place module
	repo := children_item_repository.NewChildrenItemRepository(db)
	service := children_item_service.NewChildrenItemService(repo)
	http_handler := children_item_http.NewCategoryHandler(service)

	r.POST("/children-item", http_handler.Create())
	r.GET("/children-item", http_handler.FindAll())
	r.GET("/children-item/:id", http_handler.FindOneByID())
	r.PUT("/children-item/:id", http_handler.Update())
	r.DELETE("/children-item/:id", http_handler.Delete())

}
