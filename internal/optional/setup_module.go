package item_optional_module

import (
	children_item_repository "Food-Delivery/internal/children_item/repository"
	children_item_service "Food-Delivery/internal/children_item/service"
	item_optional_http "Food-Delivery/internal/optional/controller/http"
	item_optional_repository "Food-Delivery/internal/optional/repository"
	item_optional_service "Food-Delivery/internal/optional/service"
	"Food-Delivery/pkg/app_context"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {
	db := appCtx.GetDbContext().GetMainConnection()

	childrenItem_repo := children_item_repository.NewChildrenItemRepository(db)
	childrenItem_service := children_item_service.NewChildrenItemService(childrenItem_repo)

	//dependency of place module
	repo := item_optional_repository.NewItemOptionalRepository(db)
	service := item_optional_service.NewItemOptionalService(repo, childrenItem_service)
	http_handler := item_optional_http.NewItemOptionalHandler(service)

	r.GET("/item-optional", http_handler.FindAll())
	r.GET("/item-optional/:id", http_handler.FindOneByID())
	r.POST("/item-optional", http_handler.Create())
	r.PUT("/item-optional/:id", http_handler.Update())
	r.DELETE("/item-optional/:id", http_handler.Delete())

}
