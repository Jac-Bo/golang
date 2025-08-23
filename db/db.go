package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbOpenAndUpdate() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("HOST"), os.Getenv("DBNAME"))

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("message: ", err)
	}

	db.AutoMigrate(&Article)

}
