package models

type Product struct {
	Id               int64            `gorm:"primaryKey" json:"id"`
	Name             string           `gorm:"type:varchar(50)" json:"name" binding:"required,min=3"`
	Stock            int64            `json:"stock"`
	Price            int64            `json:"price"`
	StockDecreaseLog StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
