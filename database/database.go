package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	var err error

	dsn := "root:fiangumilar@tcp(127.0.0.1:3306)/gorm_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't open database")
	}
	fmt.Println("Connect to database...")
}
