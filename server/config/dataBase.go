package config

import (
	"fmt"
	"log"
	"os"
	"server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DATA BASE ERROR")
	}
	db.AutoMigrate(&models.User{}, &models.Artis{}, &models.Song{}, &models.Favorite{})

	DB = db
	log.Println("Succes Connect DataBase")
}
