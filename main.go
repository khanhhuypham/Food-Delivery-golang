package main

import (
	"Food-Delivery/config"
	category_module "Food-Delivery/internal/category"
	order_item_module "Food-Delivery/internal/order_item"

	menu_item_module "Food-Delivery/internal/menu_item"
	"Food-Delivery/internal/middleware"
	order_module "Food-Delivery/internal/order"
	restaurant_module "Food-Delivery/internal/restaurant"
	user_module "Food-Delivery/internal/user"
	user_repository "Food-Delivery/internal/user/repository"
	"Food-Delivery/pkg/db/mysql"
	"Food-Delivery/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu_item item from here.</p>

func main() {
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
	r.Use(middlewareManager.Recover())
	v1 := r.Group("/api/v1")

	category_module.Setup(db, v1)
	restaurant_module.Setup(db, v1)
	menu_item_module.Setup(db, v1)
	order_module.Setup(db, v1)
	order_item_module.Setup(db, v1)
	user_module.Setup(db, v1, cfg, hasher, middlewareManager)

	r.Run(fmt.Sprintf(":%s", cfg.App.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
