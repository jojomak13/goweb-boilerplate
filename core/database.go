package core

import (
	"fmt"
	"log"

	"github.com/jojomak13/goweb-biolerplate/config"
	"github.com/jojomak13/goweb-biolerplate/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(c config.AppConfig) {
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", 
		c.DBUser, 
		c.DBPassword, 
		c.DBPort, 
		c.DBDataBase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("Cannont connect to database [%v]", err)
	}

	fmt.Println("Connection open to database :)")
	
	db.AutoMigrate(&model.User{})

	fmt.Println("Database migrated.")

	DB = db 
}