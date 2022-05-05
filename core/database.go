package core

import (
	"fmt"
	"log"

	"github.com/jojomak13/goweb-biolerplate/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("Cannont connect to database [%v]", err)
	}

	fmt.Println("Connection open to database :)")
	
	db.AutoMigrate(&model.User{})

	fmt.Println("Database migrated.")

	DB = db 
}