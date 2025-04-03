package main

import (
	"Food-Delivery/config"
	categorymodule "Food-Delivery/internal/category"
	"Food-Delivery/internal/middleware"
	restaurant_module "Food-Delivery/internal/restaurant"
	user_repository "Food-Delivery/internal/user/repository"
	"Food-Delivery/pkg/db/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

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

	middleware := middleware.NewMiddlewareManager(cfg, user_repository.NewUserRepository(db))
	r := gin.Default()
	r.Use(middleware.Recover())
	v1 := r.Group("/api/v1")

	categorymodule.SetupCategoryModule(db, v1)
	restaurant_module.SetupRestaurantModule(db, v1)
	r.Run(fmt.Sprintf(":%s", cfg.App.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
