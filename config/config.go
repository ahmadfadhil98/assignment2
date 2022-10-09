package config

import (
	"assignment2/structs"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=rosie.db.elephantsql.com "+
			"port=5432 "+
			"user=tnptuijq "+
			"dbname=tnptuijq "+
			"password=NVl9g3XJWZg3bX5tlYkqm9C44EdKdGyI sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(structs.Items{}, structs.Orders{})
	return db
}
