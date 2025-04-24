package media_module

import (
	"Food-Delivery/config"
	media_grpc_server "Food-Delivery/internal/media/controller/grpc-server"
	media_repository "Food-Delivery/internal/media/repository"
	media_service "Food-Delivery/internal/media/service"
	"Food-Delivery/pkg/upload"
	"Food-Delivery/proto-buffer/gen/mediapb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

func Setup(db *gorm.DB, r *gin.RouterGroup, cfg *config.Config) {

	//Declare s3
	s3Provider := upload.NewS3Provider(cfg)

	//Declare service
	repo := media_repository.NewMediaRepository(db)
	service := media_service.NewMediaService(repo, s3Provider)

	//handler := media_http.NewMediaHandler(s3Provider, service)
	//r.POST("/media/upload", handler.Upload())

	//run grpc-server server
	go func() {
		//Create a listener on TPC port
		listen, err := net.Listen("tcp", cfg.Grpc.Url)
		if err != nil {
			log.Fatalln("Failed to listen:", err)
		}
		// Create a gRPC server object
		s := grpc.NewServer()
		// Attach the Greeter service to the server
		mediapb.RegisterMediaServer(s, media_grpc_server.NewMediaGrpcServer(s3Provider, service))
		// Serve gRPC Server

		log.Printf("media_module: Serving gRPC on %v", cfg.Grpc.Url)

		err = s.Serve(listen)

		if err != nil {
			log.Fatal("error while serving on localhost:5006")
		}
	}()
}
