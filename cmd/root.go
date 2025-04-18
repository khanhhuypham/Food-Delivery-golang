package cmd

import (
	"Food-Delivery/config"
	"Food-Delivery/entity/model"
	category_module "Food-Delivery/internal/category"
	item_module "Food-Delivery/internal/item"
	upload_module "Food-Delivery/internal/media"
	"Food-Delivery/internal/middleware"
	order_module "Food-Delivery/internal/order"
	order_item_module "Food-Delivery/internal/order_item"
	restaurant_module "Food-Delivery/internal/restaurant"
	user_module "Food-Delivery/internal/user"
	user_repository "Food-Delivery/internal/user/repository"
	"Food-Delivery/pkg/db/mysql"
	"Food-Delivery/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"os"
)

//var consumerCmd = &cobra.Command{
//	Use:   "consumer",
//	Short: "Start consumer",
//}

var rootCmd = &cobra.Command{
	Use:   "increase-like",
	Short: "Start consumer increase like count restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		mainSetup()
	},
}

//var consumerDecreaseLikeCountRestaurantCmd = &cobra.Command{
//	Use:   "decrease-like",
//	Short: "Start consumer decrease like count restaurant",
//	Run: func(cmd *cobra.Command, args []string) {
//		log.Println("Start consumer decrease like count restaurant")
//	},
//}

func Execute() {
	setupConsumerCmd()
	rootCmd.AddCommand(consumerCmd)
	rootCmd.AddCommand(testNatsCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Failed to execute command;", err)
	}
}

func mainSetup() {
	env := os.Getenv("ENV")
	fileName := "./config/config-local.yml"

	fmt.Println("ENV:", env)
	//if strings.ToLower(env) == "development" {
	//	fileName = "./config/config-development.yml"
	//}
	cfg, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalln("db connection err: ", err)
	}
	db, err := mysql.MySQLConnection(cfg)

	if err != nil {
		log.Fatal("Cannot connect mysql: ", err)
		return
	}

	middlewareManager := middleware.NewMiddlewareManager(cfg, user_repository.NewUserRepository(db))
	hasher := utils.NewHashIds(cfg.App.Secret, 10)

	r := gin.Default()
	r.Use(
		middlewareManager.Recover(),
		middlewareManager.CORS(),
	)
	v1 := r.Group("/api/v1")
	model.SetupModel(db)
	category_module.Setup(db, v1, cfg)
	restaurant_module.Setup(db, v1)
	item_module.Setup(db, v1)
	order_module.Setup(db, v1)
	order_item_module.Setup(db, v1)
	user_module.Setup(db, v1, cfg, hasher, middlewareManager)
	upload_module.Setup(db, v1, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.App.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
