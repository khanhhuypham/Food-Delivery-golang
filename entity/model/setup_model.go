package model

import (
	"gorm.io/gorm"
	"log"
)

func SetupModel(db *gorm.DB) {
	/*
		remember to allign these below model in a correct order
	*/
	if err := db.AutoMigrate(
		&Category{},
		&Restaurant{},
		&VendorCategory{},
		&User{},
		&Item{},
		&Driver{},
		&Order{},
		&Rating{},
		&ChildrenItem{},
		&Optional{},
		//&OrderItem{},
		//&ItemOnChildrenItems{},
	); err != nil {
		log.Fatalf("could not migrate schema: %v", err)
	}

}
