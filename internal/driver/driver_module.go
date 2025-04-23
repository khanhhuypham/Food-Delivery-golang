package driver_module

import (
	driver_http_handler "Food-Delivery/internal/driver/controller/http"
	driver_repository "Food-Delivery/internal/driver/repository"
	driver_service "Food-Delivery/internal/driver/service"
	"Food-Delivery/pkg/app_context"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {
	db := appCtx.GetDbContext().GetMainConnection()
	//_ := appCtx.GetConfig()

	//dependency of place module
	repo := driver_repository.NewDriverRepository(db)
	service := driver_service.NewDriverService(repo)
	http_handler := driver_http_handler.NewDriverHandler(service)

	r.POST("/driver", http_handler.Create())
	r.GET("/driver", http_handler.FindAll())
	r.GET("/driver/:id", http_handler.FindOneByID())
	r.PUT("/driver/:id", http_handler.Update())
	r.PATCH("/driver/:id", http_handler.Update())
	r.DELETE("/driver/:id", http_handler.Delete())

}
