package upload_module

import (
	media_grpc_client "Food-Delivery/internal/upload/controller/grpc-client"
	upload_http_handler "Food-Delivery/internal/upload/controller/http"
	"Food-Delivery/pkg/app_context"
	"github.com/gin-gonic/gin"
)

func Setup(appCtx app_context.AppContext, r *gin.RouterGroup) {

	cfg := appCtx.GetConfig()
	//dependency of place module
	mediaGRPCClient := media_grpc_client.NewMediaGRPCClient(cfg.Grpc.Url)
	handler := upload_http_handler.NewUploadHandler(mediaGRPCClient)
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

	r.POST("/upload", handler.Upload())

}
