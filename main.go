package main

import "Food-Delivery/cmd"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> item item from here.</p>
//
//func main() {
//	env := os.Getenv("ENV")
//	fileName := "./config/config-local.yml"
//
//	fmt.Println("ENV:", env)
//	//if strings.ToLower(env) == "development" {
//	//	fileName = "./config/config-development.yml"
//	//}
//	cfg, err := config.LoadConfig(fileName)
//	if err != nil {
//		log.Fatalln("db connection err: ", err)
//	}
//	db, err := mysql.MySQLConnection(cfg)
//
//	if err != nil {
//		log.Fatal("Cannot connect mysql: ", err)
//		return
//	}
//
//	middlewareManager := middleware.NewMiddlewareManager(cfg, user_repository.NewUserRepository(db))
//	hasher := utils.NewHashIds(cfg.App.Secret, 10)
//
//	r := gin.Default()
//	r.Use(
//		middlewareManager.Recover(),
//		middlewareManager.CORS(),
//	)
//	v1 := r.Group("/api/v1")
//	model.SetupModel(db)
//	category_module.Setup(db, v1, cfg)
//	restaurant_module.Setup(db, v1)
//	item_module.Setup(db, v1)
//	order_module.Setup(db, v1)
//	order_item_module.Setup(db, v1)
//	user_module.Setup(db, v1, cfg, hasher, middlewareManager)
//	upload_module.Setup(db, v1, cfg)
//
//	r.Run(fmt.Sprintf(":%s", cfg.App.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}

func main() {
	cmd.Execute()
}
