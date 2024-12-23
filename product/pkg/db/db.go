package db

import (
	"log"

	"github.com/kwul0208/product/pkg/config"
	models "github.com/kwul0208/product/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(config config.Config) Handler {
	db, err := gorm.Open(mysql.Open(config.DBUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed init DB", err)
	}
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.StockDecreaseLog{})
	return Handler{db}
}
