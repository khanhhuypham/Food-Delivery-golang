package mysql

import (
	"Food-Delivery/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MySQLConnection(config *config.Config) (*gorm.DB, error) {
	cfg := config.Mysql
	user := cfg.User
	password := cfg.Password
	host := cfg.Host
	port := cfg.Port
	DBName := cfg.DbName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
		return nil, err
	}

	return db, nil
}
