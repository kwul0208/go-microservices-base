package initializers

import (
	"os"

	models "github.com/kwul0208/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dns := os.Getenv("DB")
	database, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{})

	DB = database

	return database
}
