package rating_module

import (
	rating_http_handler "Food-Delivery/internal/rating/controller/http"
	rating_repository "Food-Delivery/internal/rating/repository"
	rating_service "Food-Delivery/internal/rating/service"
	"Food-Delivery/pkg/app_context"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {
	db := appCtx.GetDbContext().GetMainConnection()
	repo := rating_repository.NewRatingRepository(db)
	service := rating_service.NewRatingService(repo, appCtx.GetMsgBroker())
	handler := rating_http_handler.NewRatingHandler(service)

	r.GET("/rating/like", handler.Like())
	r.POST("/rating/comment", handler.Comment())
	r.POST("/rating/score", handler.SetScore())
}
