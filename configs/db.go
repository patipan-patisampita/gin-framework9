package configs

import (
	"fmt"
	"os"

	"github.com/patipan-patisampita/gin-framework9/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DATABASES_DNS")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Not connect Databases")
		fmt.Println(err.Error())
	}
	fmt.Println("Success connction Databases")

	//Migration
	db.AutoMigrate(&models.User{},&models.Blog{})

	DB = db
}
