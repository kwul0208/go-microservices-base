package initializers

import (
	"log"
	"os"

	models "github.com/kwul0208/product/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dns := os.Getenv("DB")
	database, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		// panic(err)
		log.Println(err)
	}

	database.AutoMigrate(&models.Product{})

	DB = database

	return database
}
